package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

type cronJobController struct{}

var CronJobController = &cronJobController{}

func (c *cronJobController) Add(ctx context.Context, req *v1.AddCronJobReq) (res *v1.AddCronJobRes, err error) {
	id, err := service.CronJob().Add(ctx, req)
	if err != nil {
		return
	}

	res = &v1.AddCronJobRes{
		ID: id,
	}
	return
}

func (c *cronJobController) Delete(ctx context.Context, req *v1.DeleteCronJobReq) (res *v1.DeleteCronJobRes, err error) {
	err = service.CronJob().Delete(ctx, req.ID)
	if err != nil {
		return
	}
	return
}

func (c *cronJobController) Edit(ctx context.Context, req *v1.EditCronJobReq) (res *v1.EditCronJobRes, err error) {
	err = service.CronJob().Edit(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *cronJobController) Get(ctx context.Context, req *v1.GetCronJobReq) (res *v1.GetCronJobRes, err error) {
	out, err := service.CronJob().Get(ctx, req.ID)
	if err != nil {
		return
	}

	res = &v1.GetCronJobRes{
		CronJob: &v1.CronJob{
			ID:             out.ID,
			Name:           out.Name,
			Params:         out.Params,
			InvokeTarget:   out.InvokeTarget,
			CronExpression: out.CronExpression,
			Enabled:        out.Enabled == model.CronJobStatusEnabled,
			LastExecuteAt:  out.LastExecuteAt,
			NextExecuteAt:  out.NextExecuteAt,
			ExecuteCount:   out.ExecuteCount,
			SuccessCount:   out.SuccessCount,
			FailedCount:    out.FailedCount,
			Remark:         out.Remark,
		},
	}
	return
}

func (c *cronJobController) GetList(ctx context.Context, req *v1.GetCronJobListReq) (res *v1.GetCronJobListRes, err error) {
	out, total, err := service.CronJob().GetList(ctx, req)
	if err != nil {
		return
	}

	res = &v1.GetCronJobListRes{
		PageRes: model.PageRes{
			Total: total,
		},
	}
	for _, item := range out {
		res.List = append(res.List, &v1.CronJob{
			ID:             item.ID,
			Name:           item.Name,
			Params:         item.Params,
			InvokeTarget:   item.InvokeTarget,
			CronExpression: item.CronExpression,
			Enabled:        item.Enabled == model.CronJobStatusEnabled,
			LastExecuteAt:  item.LastExecuteAt,
			NextExecuteAt:  item.NextExecuteAt,
			ExecuteCount:   item.ExecuteCount,
			SuccessCount:   item.SuccessCount,
			FailedCount:    item.FailedCount,
			Remark:         item.Remark,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
		})
	}

	return
}

// GetJobLogs 获取任务执行日志
func (c *cronJobController) GetJobLogs(ctx context.Context, req *v1.GetJobLogsReq) (res *v1.GetJobLogsRes, err error) {
	logs, total, err := service.CronJob().GetJobLogs(ctx, req)
	if err != nil {
		return
	}

	res = &v1.GetJobLogsRes{
		PageRes: model.PageRes{
			Total: total,
		},
	}

	for _, log := range logs {
		res.List = append(res.List, &v1.CronJobLog{
			ID:        log.ID,
			JobID:     log.JobID,
			Status:    int(log.ExecuteStatus),
			StartTime: log.StartTime,
			EndTime:   log.EndTime,
			Duration:  log.Duration,
			CreatedAt: log.CreatedAt,
		})
	}

	return
}
