package service

import (
	"DeviceManagement/internal/model"
	"context"
)

type IDeviceLogService interface {
	Add(ctx context.Context, in *model.DeviceOnlineLogReq) (err error)

	List(ctx context.Context, req *model.DeviceOnlineLogListReq) (out []*model.DeviceOnlineLog, pageRes *model.PageRes, err error)
}

var localDeviceLog IDeviceLogService

func DeviceLog() IDeviceLogService {
	if localDeviceLog == nil {
		panic("implement not found for interface IDeviceLogService, forgot register?")
	}
	return localDeviceLog
}

func RegisterDeviceLog(i IDeviceLogService) {
	localDeviceLog = i
}
