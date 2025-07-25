package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IDeviceLog interface {
	List(ctx context.Context, req *v1.DeviceLogListReq) (out []*model.DeviceLog, pageRes *model.PageRes, err error)
}

var localDeviceLog IDeviceLog

func DeviceLog() IDeviceLog {
	if localDeviceLog == nil {
		panic("implement not found for interface IDeviceLog, forgot register?")
	}
	return localDeviceLog
}

func RegisterDeviceLog(i IDeviceLog) {
	localDeviceLog = i
}
