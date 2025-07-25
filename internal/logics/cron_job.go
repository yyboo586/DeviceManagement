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

	"github.com/gogf/gf/v2/frame/g"
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

type DeviceCountResult struct {
	Total int `json:"total"`
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
		g.Log().Error(ctx, err)
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
		g.Log().Error(ctx, err)
		return
	}

	schedulerInstance.Notify(ctx, id)

	return
}

func (c *cronJob) Delete(ctx context.Context, id string) (err error) {
	_, err = dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().ID, id).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (c *cronJob) Edit(ctx context.Context, in *v1.EditCronJobReq) (err error) {
	updateData := map[string]interface{}{
		dao.CronJob.Columns().Name:           in.Name,
		dao.CronJob.Columns().CronExpression: in.CronExpression,
	}

	result, err := dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().ID, in.ID).Data(updateData).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if rowsAffected == 0 {
		err = fmt.Errorf("任务不存在")
		g.Log().Error(ctx, err)
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
		g.Log().Error(ctx, err)
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
		g.Log().Error(ctx, err)
		return
	}

	entities := make([]*entity.TCronJob, 0)
	err = m.Page(in.PageReq.Page, in.PageReq.PageSize).Scan(&entities)
	if err != nil {
		g.Log().Error(ctx, err)
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
		g.Log().Error(ctx, err)
		return
	}

	for _, entity := range entities {
		out = append(out, c.convertToModel(entity))
	}
	return
}

func (c *cronJob) ExecuteJob(ctx context.Context, jobID string) (err error) {
	jobInfo, err := c.Get(ctx, jobID)
	if err != nil {
		return
	}

	schedulerInstance.AddTask(ctx, jobInfo)

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
