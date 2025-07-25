package service

import (
	"DeviceManagement/internal/model"
	"context"
)

type IDeviceService interface {
	Add(ctx context.Context, device *model.Device) (id int64, err error)

	Delete(ctx context.Context, ids []int64) (err error)

	Update(ctx context.Context, device *model.Device) (err error)
	EditDeviceStatus(ctx context.Context, ids []int64, enabled bool) (err error)

	Get(ctx context.Context, id int64) (out *model.Device, err error)
	List(ctx context.Context, orgID string, page *model.PageReq) (list []*model.Device, pageRes *model.PageRes, err error)

	// 设备告警
	InvokeAlarm(ctx context.Context, orgID string, content string) (err error)

	DeviceCountHandler(ctx context.Context, params interface{}) (success bool, result string)
}

var localDevice IDeviceService

func Device() IDeviceService {
	if localDevice == nil {
		panic("implement not found for interface IDeviceService, forgot register?")
	}
	return localDevice
}

func RegisterDevice(i IDeviceService) {
	localDevice = i
}
