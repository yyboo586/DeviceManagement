package logics

import (
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

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
	cron *gcron.Cron
	jobs map[string]*model.CronJob // 任务ID -> 任务信息
	mu   sync.RWMutex

	ch chan string
}

func NewScheduler() *scheduler {
	schedulerOnce.Do(func() {
		schedulerInstance = &scheduler{
			cron: gcron.New(),
			jobs: make(map[string]*model.CronJob),
			ch:   make(chan string, 100),
		}

		go schedulerInstance.worker(context.Background())
	})
	return schedulerInstance
}

func (s *scheduler) Start(ctx context.Context) error {
	// 加载所有启用的任务
	jobs, err := cronJobInstance.GetEnabledJobs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get enabled jobs: %w", err)
	}

	// 添加所有任务到调度器
	for _, job := range jobs {
		if err := s.AddJob(ctx, job); err != nil {
			g.Log().Errorf(ctx, "failed to add job %s: %v", job.ID, err)
			continue
		}
	}

	// 启动调度器
	s.cron.Start()
	g.Log().Info(ctx, "cron job scheduler started")
	return nil
}

// Stop 停止调度器
func (s *scheduler) Stop() {
	s.cron.Stop()
	g.Log().Info(context.Background(), "cron job scheduler stopped")
}

func (s *scheduler) Notify(ctx context.Context, jobID string) error {
	select {
	case s.ch <- jobID:
	default:
		return fmt.Errorf("channel is full")
	}
	return nil
}

func (s *scheduler) worker(ctx context.Context) {
	for jobID := range s.ch {
		jobInfo, err := cronJobInstance.Get(ctx, jobID)
		if err != nil {
			g.Log().Errorf(ctx, "failed to get job %s: %v", jobID, err)
			continue
		}
		err = s.AddJob(ctx, jobInfo)
		if err != nil {
			g.Log().Errorf(ctx, "failed to add job %s: %v", jobID, err)
		}
	}
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
	entry, err := s.cron.Add(ctx, jobInfo.CronExpression, s.createJobFunc(ctx, jobInfo), fmt.Sprintf("job-%s", jobInfo.ID))
	if err != nil {
		return fmt.Errorf("failed to add cron job: %w", err)
	}

	jobInfo.Entry = entry
	s.jobs[jobInfo.ID] = jobInfo

	g.Log().Infof(ctx, "added job %s (%s) with cron expression: %s", jobInfo.ID, jobInfo.Name, jobInfo.CronExpression)
	return nil
}

// RemoveJob 从调度器移除任务
func (s *scheduler) RemoveJob(jobID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.removeJob(jobID)
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

// createJobFunc 创建任务执行函数
func (s *scheduler) createJobFunc(jobCtx context.Context, jobInfo *model.CronJob) gcron.JobFunc {
	return func(ctx context.Context) {
		startTime := gtime.Now()

		// 新增任务运行记录
		var jobLogInfo *model.CronJobLog = &model.CronJobLog{
			OrgID:     jobInfo.OrgID,
			JobID:     jobInfo.ID,
			StartTime: startTime,
		}

		jobLogID, err := service.CronJob().AddJobLog(jobCtx, jobLogInfo)
		if err != nil {
			g.Log().Errorf(jobCtx, "failed to add job log: %v", err)
		}
		jobLogInfo.ID = jobLogID

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
		result, err := json.Marshal(message)
		if err != nil {
			g.Log().Errorf(jobCtx, "failed to marshal message: %v", err)
			return
		}

		endTime := gtime.Now()
		duration := endTime.Sub(startTime).Milliseconds()

		// 计算下次执行时间
		nextExecuteAt := s.calculateNextExecuteTime(jobInfo.CronExpression, endTime)

		jobStats := map[string]interface{}{
			dao.CronJob.Columns().ID:            jobInfo.ID,
			dao.CronJob.Columns().LastExecuteAt: endTime,
			dao.CronJob.Columns().NextExecuteAt: nextExecuteAt,
			dao.CronJob.Columns().ExecuteCount:  jobInfo.ExecuteCount + 1,
		}
		jobLogStats := map[string]interface{}{
			dao.CronJobLog.Columns().Result:   string(result),
			dao.CronJobLog.Columns().EndTime:  endTime,
			dao.CronJobLog.Columns().Duration: duration,
		}
		if success {
			jobStats[dao.CronJob.Columns().LastExecuteStatus] = model.CronJobExecuteStatusSuccess
			jobStats[dao.CronJob.Columns().SuccessCount] = jobInfo.SuccessCount + 1

			jobLogStats[dao.CronJobLog.Columns().ExecuteStatus] = model.CronJobExecuteStatusSuccess
		} else {
			jobStats[dao.CronJob.Columns().LastExecuteStatus] = model.CronJobExecuteStatusFailed
			jobStats[dao.CronJob.Columns().FailedCount] = jobInfo.FailedCount + 1

			jobLogStats[dao.CronJobLog.Columns().ExecuteStatus] = model.CronJobExecuteStatusFailed
		}

		err = g.DB().Transaction(jobCtx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.CronJob.Ctx(ctx).TX(tx).Data(jobStats).Where(dao.CronJob.Columns().ID, jobInfo.ID).Update()
			if err != nil {
				return err
			}
			_, err = dao.CronJobLog.Ctx(ctx).TX(tx).Data(jobLogStats).Where(dao.CronJobLog.Columns().ID, jobLogInfo.ID).Update()
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			g.Log().Errorf(jobCtx, "failed to update job stats: %v", err)
		}

		g.Log().Infof(jobCtx, "job %s (%s) executed, success: %v, duration: %dms, message: %s",
			jobInfo.ID, jobInfo.Name, success, duration, message)
	}
}

// executeJob 执行具体的任务
func (s *scheduler) executeJob(ctx context.Context, jobInfo *model.CronJob) (success bool, result string) {
	// 获取执行器工厂
	factory := NewExecutorFactory()

	// 执行任务
	out := factory.Execute(ctx, ExecutorTypeHTTP, jobInfo.Params)

	// 记录详细的执行结果
	if out.Error != nil {
		g.Log().Errorf(ctx, "job execution error: %v", out.Error)
	}

	g.Log().Debugf(ctx, "job execution result: success=%v, duration=%dms, message=%s",
		out.Success, out.Duration, out.Message)

	return out.Success, out.Message
}

// calculateNextExecuteTime 计算下次执行时间
func (s *scheduler) calculateNextExecuteTime(cronExpression string, currentTime *gtime.Time) *gtime.Time {
	// 这里可以使用cron库来计算下次执行时间
	// 简化实现，实际可以使用robfig/cron等库
	return currentTime.Add(time.Minute) // 示例：1分钟后
}

// GetJobInfo 获取任务信息
func (s *scheduler) GetJobInfo(jobID string) *model.CronJob {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.jobs[jobID]
}

// GetAllJobs 获取所有任务信息
func (s *scheduler) GetAllJobs() map[string]*model.CronJob {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]*model.CronJob)
	for k, v := range s.jobs {
		result[k] = v
	}
	return result
}
