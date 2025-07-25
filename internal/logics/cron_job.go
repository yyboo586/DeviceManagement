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

func (c *cronJob) Add(ctx context.Context, in *v1.AddCronJobReq, cronExpression, description string) (id string, err error) {
	template, err := cronJobTemplateInstance.Get(ctx, in.TemplateID)
	if err != nil {
		return
	}
	if template == nil {
		err = fmt.Errorf("任务模板不存在")
		return
	}

	id = common.GetSnowflakeID()
	insertData := map[string]interface{}{
		dao.CronJob.Columns().ID:             id,
		dao.CronJob.Columns().OrgID:          in.OrgID,
		dao.CronJob.Columns().TemplateID:     template.ID,
		dao.CronJob.Columns().Name:           template.Name,
		dao.CronJob.Columns().Enabled:        model.CronJobStatusEnabled,
		dao.CronJob.Columns().Remark:         in.Remark,
		dao.CronJob.Columns().Params:         template.Config, // 占位，待后续支持
		dao.CronJob.Columns().InvokeType:     template.InvokeType,
		dao.CronJob.Columns().CronExpression: cronExpression,
		dao.CronJob.Columns().Description:    description,
	}

	_, err = dao.CronJob.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = fmt.Errorf("任务已存在")
		}
		g.Log().Error(ctx, err)
		return
	}

	// 将新任务添加到调度器
	schedulerInstance.AddJob(ctx, &model.CronJob{
		ID:                id,
		OrgID:             in.OrgID,
		Name:              template.Name,
		Enabled:           model.CronJobStatusEnabled,
		Remark:            "",
		Params:            template.Config,
		InvokeType:        template.InvokeType,
		CronExpression:    cronExpression,
		LastExecuteStatus: model.CronJobExecuteStatus(0), // 等待执行状态
		ExecuteCount:      0,
		SuccessCount:      0,
		FailedCount:       0,
	})

	return
}

func (c *cronJob) Delete(ctx context.Context, id string) (err error) {
	_, err = dao.CronJob.Ctx(ctx).Where(dao.CronJob.Columns().ID, id).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 从调度器移除任务
	schedulerInstance.RemoveJob(id)

	return
}

func (c *cronJob) Edit(ctx context.Context, in *v1.EditCronJobReq, cronExpression, description string) (err error) {
	updateData := map[string]interface{}{
		dao.CronJob.Columns().Name:           in.Name,
		dao.CronJob.Columns().CronExpression: cronExpression,
		dao.CronJob.Columns().Description:    description,
	}
	if in.Remark != "" {
		updateData[dao.CronJob.Columns().Remark] = in.Remark
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

	// 更新调度器中的任务
	schedulerInstance.UpdateJob(ctx, in.ID)

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
	// 直接调用调度器的手动执行方法
	return schedulerInstance.ExecuteJobNow(ctx, jobID)
}

func (c *cronJob) convertToModel(in *entity.TCronJob) *model.CronJob {
	return &model.CronJob{
		ID:                in.ID,
		OrgID:             in.OrgID,
		Name:              in.Name,
		Enabled:           model.CronJobStatus(in.Enabled),
		Remark:            in.Remark,
		Params:            in.Params,
		InvokeType:        in.InvokeType,
		CronExpression:    in.CronExpression,
		Description:       in.Description,
		LastExecuteStatus: model.CronJobExecuteStatus(in.LastExecuteStatus),
		LastExecuteAt:     in.LastExecuteAt,
		ExecuteCount:      in.ExecuteCount,
		SuccessCount:      in.SuccessCount,
		FailedCount:       in.FailedCount,
		CreatedAt:         in.CreatedAt,
		UpdatedAt:         in.UpdatedAt,
	}
}
