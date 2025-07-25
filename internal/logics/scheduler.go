package logics

import (
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	schedulerOnce     sync.Once
	schedulerInstance *scheduler
)

type scheduler struct {
	ctx      context.Context
	cron     *gcron.Cron
	jobs     map[string]*model.CronJob // 任务ID -> 任务信息
	handlers map[string]model.CronJobHandler
	mu       sync.RWMutex
}

func NewScheduler() *scheduler {
	schedulerOnce.Do(func() {
		schedulerInstance = &scheduler{
			ctx:      context.Background(),
			cron:     gcron.New(),
			jobs:     make(map[string]*model.CronJob),
			handlers: make(map[string]model.CronJobHandler),
		}
	})
	return schedulerInstance
}

// Start 启动调度器，加载所有启用的任务
func (s *scheduler) Start() error {
	jobs, err := cronJobInstance.GetEnabledJobs(s.ctx)
	if err != nil {
		return fmt.Errorf("failed to get enabled jobs: %w", err)
	}

	for _, job := range jobs {
		if err := s.AddJob(s.ctx, job); err != nil {
			g.Log().Errorf(s.ctx, "failed to add job %s: %v", job.ID, err)
			continue
		}
	}

	// 启动调度器
	s.cron.Start()
	g.Log().Info(s.ctx, "cron job scheduler started")
	return nil
}

// Stop 停止调度器
func (s *scheduler) Stop() {
	s.cron.Stop()
	g.Log().Info(s.ctx, "cron job scheduler stopped")
}

// AddJob 添加任务到调度器
func (s *scheduler) AddJob(ctx context.Context, jobInfo *model.CronJob) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 如果任务已存在，先移除
	if _, exists := s.jobs[jobInfo.ID]; exists {
		s.removeJob(jobInfo.ID)
	}

	// 添加任务到cron调度器
	entry, err := s.cron.Add(s.ctx, jobInfo.CronExpression, s.createJobFunc(s.ctx, jobInfo.ID), fmt.Sprintf("job-%s", jobInfo.ID))
	if err != nil {
		return fmt.Errorf("failed to add cron job: %w", err)
	}

	jobInfo.Entry = entry
	s.jobs[jobInfo.ID] = jobInfo

	g.Log().Infof(s.ctx, "added job %s (%s) with cron expression: %s", jobInfo.ID, jobInfo.Name, jobInfo.CronExpression)
	return nil
}

// RemoveJob 从调度器移除任务
func (s *scheduler) RemoveJob(jobID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.removeJob(jobID)
}

// UpdateJob 更新任务（先移除再添加）
func (s *scheduler) UpdateJob(ctx context.Context, jobID string) error {
	// 先从调度器移除
	s.RemoveJob(jobID)

	// 从数据库重新获取任务信息
	jobInfo, err := cronJobInstance.Get(ctx, jobID)
	if err != nil {
		return fmt.Errorf("failed to get job %s: %w", jobID, err)
	}

	// 如果任务仍然启用，重新添加到调度器
	if jobInfo.Enabled == model.CronJobStatusEnabled {
		return s.AddJob(s.ctx, jobInfo)
	}

	return nil
}

// ExecuteJobNow 立即执行一次任务（手动触发）
func (s *scheduler) ExecuteJobNow(ctx context.Context, jobID string) error {
	jobInfo, err := cronJobInstance.Get(ctx, jobID)
	if err != nil {
		return fmt.Errorf("failed to get job %s: %w", jobID, err)
	}

	token := ctx.Value(common.BearerToken).(string)
	// 立即执行任务
	s.createJobFunc(context.WithValue(ctx, common.BearerToken, token), jobInfo.ID)(s.ctx)
	return nil
}

func (s *scheduler) removeJob(jobID string) {
	if jobInfo, exists := s.jobs[jobID]; exists {
		if jobInfo.Entry != nil {
			s.cron.Remove(fmt.Sprintf("job-%s", jobID))
		}
		delete(s.jobs, jobID)
		g.Log().Infof(context.Background(), "removed job %s", jobID)
	}
}

func (s *scheduler) getCronJob(jobID string) (jobInfo *model.CronJob, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	jobInfo, exists := s.jobs[jobID]
	if !exists {
		return nil, fmt.Errorf("job %s not found", jobID)
	}
	return
}

func (s *scheduler) setCronJob(jobID string, jobInfo *model.CronJob) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.jobs[jobID] = jobInfo
}

// createJobFunc 创建任务执行函数
func (s *scheduler) createJobFunc(jobCtx context.Context, jobID string) gcron.JobFunc {
	return func(ctx context.Context) {
		startTime := gtime.Now()

		jobInfo, err := s.getCronJob(jobID)
		if err != nil {
			g.Log().Errorf(ctx, "failed to get job %s: %v", jobID, err)
			return
		}
		defer func() {
			s.setCronJob(jobInfo.ID, jobInfo)
		}()

		insertData := map[string]interface{}{
			dao.CronJobLog.Columns().OrgID:         jobInfo.OrgID,
			dao.CronJobLog.Columns().JobID:         jobInfo.ID,
			dao.CronJobLog.Columns().JobName:       jobInfo.Name,
			dao.CronJobLog.Columns().ExecuteStatus: model.CronJobExecuteStatusRunning,
			dao.CronJobLog.Columns().StartTime:     startTime,
		}

		jobLogID, err := dao.CronJobLog.Ctx(ctx).Data(insertData).InsertAndGetId()
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}

		var success bool
		var message string

		defer func() {
			if r := recover(); r != nil {
				message = fmt.Sprintf("任务执行异常: %v", r)
				success = false
			}
		}()

		// 这里可以扩展为调用不同的目标类型
		success, message = s.executeJob(jobCtx, jobInfo)

		endTime := gtime.Now()
		jobInfo.ExecuteCount = jobInfo.ExecuteCount + 1
		if success {
			jobInfo.LastExecuteStatus = model.CronJobExecuteStatusSuccess
			jobInfo.SuccessCount = jobInfo.SuccessCount + 1
		} else {
			jobInfo.LastExecuteStatus = model.CronJobExecuteStatusFailed
			jobInfo.FailedCount = jobInfo.FailedCount + 1
		}

		jobStats := map[string]interface{}{
			dao.CronJob.Columns().LastExecuteStatus: jobInfo.LastExecuteStatus,
			dao.CronJob.Columns().LastExecuteAt:     startTime,
			dao.CronJob.Columns().ExecuteCount:      jobInfo.ExecuteCount,
			dao.CronJob.Columns().SuccessCount:      jobInfo.SuccessCount,
			dao.CronJob.Columns().FailedCount:       jobInfo.FailedCount,
		}
		jobLogStats := map[string]interface{}{
			dao.CronJobLog.Columns().ExecuteStatus: jobInfo.LastExecuteStatus,
			dao.CronJobLog.Columns().Result:        message,
			dao.CronJobLog.Columns().EndTime:       endTime,
			dao.CronJobLog.Columns().Duration:      endTime.Sub(startTime).Milliseconds(),
		}
		err = g.DB().Transaction(jobCtx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.CronJob.Ctx(ctx).TX(tx).Data(jobStats).Where(dao.CronJob.Columns().ID, jobInfo.ID).Update()
			if err != nil {
				return err
			}
			_, err = dao.CronJobLog.Ctx(ctx).TX(tx).Data(jobLogStats).Where(dao.CronJobLog.Columns().ID, jobLogID).Update()
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			g.Log().Errorf(jobCtx, "failed to update job stats: %v", err)
		}

		g.Log().Infof(jobCtx, "job %s (%s) executed, success: %v, duration: %dms, message: %s",
			jobInfo.ID, jobInfo.Name, success, endTime.Sub(startTime).Milliseconds(), message)
	}
}

// executeJob 执行具体的任务
func (s *scheduler) executeJob(ctx context.Context, jobInfo *model.CronJob) (success bool, result string) {
	switch jobInfo.InvokeType {
	case "http":
		var config map[string]interface{}
		err := json.Unmarshal([]byte(jobInfo.Params), &config)
		if err != nil {
			g.Log().Errorf(ctx, "failed to unmarshal job params: %v", err)
			success, result = false, "解析配置失败"
			return
		}

		// 安全地获取URL配置
		urlStr, ok := config["url"].(string)
		if !ok || urlStr == "" {
			g.Log().Errorf(ctx, "missing or invalid URL in config")
			success, result = false, "配置中缺少或无效的URL"
			return
		}

		url := fmt.Sprintf("%s?org_id=%s", urlStr, jobInfo.OrgID)
		client := common.NewHTTPClient()

		// 安全地获取Bearer Token
		bearerToken, ok := ctx.Value(common.BearerToken).(string)
		if !ok || bearerToken == "" {
			g.Log().Errorf(ctx, "missing or invalid bearer token in context")
			success, result = false, "缺少认证令牌"
			return
		}

		headers := map[string]interface{}{
			"Authorization": fmt.Sprintf("Bearer %s", bearerToken),
		}

		status, respBody, err := client.GET(ctx, url, headers)
		if err != nil {
			g.Log().Errorf(ctx, "failed to get job %s: %v", jobInfo.ID, err)
			success, result = false, "调用失败"
			return
		}
		if status != http.StatusOK {
			g.Log().Errorf(ctx, "failed to get job %s: %v", jobInfo.ID, err)
			success, result = false, "调用失败"
			return
		}

		var i interface{}
		err = json.Unmarshal(respBody, &i)
		if err != nil {
			g.Log().Errorf(ctx, "failed to unmarshal job params: %v", err)
			success, result = false, "解析响应失败"
			return
		}
		// 安全地解析响应数据
		responseMap, ok := i.(map[string]interface{})
		if !ok {
			g.Log().Errorf(ctx, "invalid response format: not a map")
			success, result = false, "响应格式无效"
			return
		}

		code, ok := responseMap["code"].(float64)
		if !ok {
			g.Log().Errorf(ctx, "invalid response code format")
			success, result = false, "响应码格式无效"
			return
		}

		if int(code) != 0 {
			message, ok := responseMap["message"].(string)
			if !ok {
				message = "未知错误"
			}
			g.Log().Errorf(ctx, "API call failed with code %d: %s", int(code), message)
			success, result = false, message
			return
		}

		// 安全地获取设备数量
		data, ok := responseMap["data"].(map[string]interface{})
		if !ok {
			g.Log().Errorf(ctx, "invalid response data format")
			success, result = false, "响应数据格式无效"
			return
		}

		total, ok := data["total"].(float64)
		if !ok {
			g.Log().Errorf(ctx, "invalid total count format")
			success, result = false, "设备数量格式无效"
			return
		}

		result = fmt.Sprintf("当前组织 %s 设备数量: %d", jobInfo.OrgID, int64(total))
		success = true
	default:
		success, result = false, "不支持的调用类型"
	}
	return
}

func (s *scheduler) RegisterHandler(jobName string, handler model.CronJobHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers[jobName] = handler
}
