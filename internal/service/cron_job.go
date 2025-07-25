package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type ICronJobService interface {
	// 添加任务
	Add(ctx context.Context, in *v1.AddCronJobReq, cronExpression, description string) (id string, err error)
	// 删除任务
	Delete(ctx context.Context, id string) error
	// 编辑任务
	Edit(ctx context.Context, in *v1.EditCronJobReq, cronExpression, description string) error
	// 获取任务
	Get(ctx context.Context, id string) (out *model.CronJob, err error)
	// 获取任务列表(不区分任务状态)
	GetList(ctx context.Context, in *v1.GetCronJobListReq) (out []*model.CronJob, total int, err error)

	// 获取所有启用的任务
	GetEnabledJobs(ctx context.Context) (out []*model.CronJob, err error)

	// 立即执行任务
	ExecuteJob(ctx context.Context, jobID string) error
}

var localCronJob ICronJobService

func CronJob() ICronJobService {
	if localCronJob == nil {
		panic("implement not found for interface ICronJobService, forgot register?")
	}
	return localCronJob
}

func RegisterCronJob(i ICronJobService) {
	localCronJob = i
}
