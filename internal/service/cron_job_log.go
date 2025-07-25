package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type ICronJobLog interface {
	// 获取任务执行日志列表
	GetLogs(ctx context.Context, in *v1.GetJobLogsReq) (out []*model.CronJobLog, currentPage int, total int, err error)
}

var localCronJobLog ICronJobLog

func CronJobLog() ICronJobLog {
	if localCronJobLog == nil {
		panic("implement not found for interface ICronJobLog, forgot register?")
	}
	return localCronJobLog
}

func RegisterCronJobLog(i ICronJobLog) {
	localCronJobLog = i
}
