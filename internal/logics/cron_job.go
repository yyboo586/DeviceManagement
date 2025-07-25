package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

var (
	cronJobOnce     sync.Once
	cronJobInstance *cronJob
)

var jobList = map[string]*model.CronJob{
	"DeviceCount": {
		Name:         "DeviceCount",
		Params:       `{"method":"GET","url":"http://127.0.0.1:9500/api/v1/device-management/devices?org_id=%s"}`,
		InvokeTarget: "http",
	},
}

type cronJob struct{}

func NewCronJob() *cronJob {
	cronJobOnce.Do(func() {
		cronJobInstance = &cronJob{}
	})
	return cronJobInstance
}

func (c *cronJob) Add(ctx context.Context, in *v1.AddCronJobReq) (id string, err error) {
	jobInfo := jobList[in.Name]
	if jobInfo == nil {
		err = fmt.Errorf("暂不支持的任务类型: %s", in.Name)
		return
	}

	id = common.GetSnowflakeID()
	insertData := map[string]interface{}{
		dao.CronJob.Columns().ID:             id,
		dao.CronJob.Columns().OrgID:          in.OrgID,
		dao.CronJob.Columns().Name:           in.Name,
		dao.CronJob.Columns().Enabled:        model.CronJobStatusEnabled,
		dao.CronJob.Columns().Remark:         "",
		dao.CronJob.Columns().Params:         "",
		dao.CronJob.Columns().InvokeTarget:   jobInfo.InvokeTarget,
		dao.CronJob.Columns().CronExpression: in.CronExpression,
	}

	_, err = dao.CronJob.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = fmt.Errorf("任务已存在")
		}
		return
	}

	schedulerInstance.Notify(ctx, id)

	return
}

func (c *cronJob) Delete(ctx context.Context, id string) (err error) {
	_, err = dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().ID, id).Delete()
	return
}

func (c *cronJob) Edit(ctx context.Context, in *v1.EditCronJobReq) (err error) {
	updateData := map[string]interface{}{
		dao.CronJob.Columns().Name:           in.Name,
		dao.CronJob.Columns().CronExpression: in.CronExpression,
	}

	result, err := dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().ID, in.ID).Data(updateData).Update()
	if err != nil {
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		err = fmt.Errorf("任务不存在")
		return
	}

	return
}

func (c *cronJob) Get(ctx context.Context, id string) (out *model.CronJob, err error) {
	entity := &entity.TCronJob{}
	err = dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().ID, id).Scan(entity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("任务不存在")
			return
		}
		return
	}

	out = c.convertToModel(entity)
	return
}

func (c *cronJob) GetList(ctx context.Context, in *v1.GetCronJobListReq) (out []*model.CronJob, total int, err error) {
	if in.PageReq.Page <= 0 {
		in.PageReq.Page = 1
	}
	if in.PageReq.PageSize <= 0 {
		in.PageReq.PageSize = 10
	}

	m := dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().OrgID, in.OrgID)

	total, err = m.Count()
	if err != nil {
		return
	}

	entities := make([]*entity.TCronJob, 0)
	err = m.Page(in.PageReq.Page, in.PageReq.PageSize).Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		out = append(out, c.convertToModel(entity))
	}
	return
}

func (c *cronJob) GetEnabledJobs(ctx context.Context) (out []*model.CronJob, err error) {
	entities := make([]*entity.TCronJob, 0)
	err = dao.CronJob.Ctx(ctx).
		Where(dao.CronJob.Columns().Enabled, model.CronJobStatusEnabled).
		Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		out = append(out, c.convertToModel(entity))
	}
	return
}

// AddJobLog 添加任务执行日志
func (c *cronJob) AddJobLog(ctx context.Context, log *model.CronJobLog) (id int64, err error) {
	insertData := map[string]interface{}{
		dao.CronJobLog.Columns().OrgID:         log.OrgID,
		dao.CronJobLog.Columns().JobID:         log.JobID,
		dao.CronJobLog.Columns().ExecuteStatus: model.CronJobExecuteStatusRunning,
		dao.CronJobLog.Columns().StartTime:     log.StartTime,
	}

	id, err = dao.CronJobLog.Ctx(ctx).Data(insertData).InsertAndGetId()
	return
}

// GetJobLogs 获取任务执行日志
func (c *cronJob) GetJobLogs(ctx context.Context, in *v1.GetJobLogsReq) (out []*model.CronJobLog, total int, err error) {
	if in.PageReq.Page <= 0 {
		in.PageReq.Page = 1
	}
	if in.PageReq.PageSize <= 0 {
		in.PageReq.PageSize = 10
	}

	m := dao.CronJobLog.Ctx(ctx).Where(dao.CronJobLog.Columns().OrgID, in.OrgID)

	if in.Name != "" {
		m = m.Where(dao.CronJobLog.Columns().JobID, in.Name)
	}

	total, err = m.Count()
	if err != nil {
		return
	}

	entities := make([]*entity.TCronJobLog, 0)
	err = m.Page(in.PageReq.Page, in.PageReq.PageSize).Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		log := &model.CronJobLog{
			ID:            entity.ID,
			JobID:         entity.JobID,
			ExecuteStatus: model.CronJobExecuteStatus(entity.ExecuteStatus),
			Result:        entity.Result,
			StartTime:     entity.StartTime,
			EndTime:       entity.EndTime,
			Duration:      entity.Duration,
			CreatedAt:     entity.CreatedAt,
		}
		out = append(out, log)
	}
	return
}

func (c *cronJob) convertToModel(in *entity.TCronJob) *model.CronJob {
	return &model.CronJob{
		ID:                in.ID,
		OrgID:             in.OrgID,
		Name:              in.Name,
		Enabled:           model.CronJobStatus(in.Enabled),
		Remark:            in.Remark,
		Params:            in.Params,
		InvokeTarget:      in.InvokeTarget,
		CronExpression:    in.CronExpression,
		LastExecuteStatus: model.CronJobExecuteStatus(in.LastExecuteStatus),
		LastExecuteAt:     in.LastExecuteAt,
		NextExecuteAt:     in.NextExecuteAt,
		ExecuteCount:      in.ExecuteCount,
		SuccessCount:      in.SuccessCount,
		FailedCount:       in.FailedCount,
		CreatedAt:         in.CreatedAt,
		UpdatedAt:         in.UpdatedAt,
	}
}
