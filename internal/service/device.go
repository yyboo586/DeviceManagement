package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IDeviceService interface {
	Add(ctx context.Context, device *model.Device) (id int64, err error)

	Delete(ctx context.Context, ids []int64) (err error)

	Update(ctx context.Context, device *model.Device) (err error)

	Get(ctx context.Context, id int64) (out *model.Device, err error)
	List(ctx context.Context, orgID string, page *model.PageReq) (list []*model.Device, pageRes *model.PageRes, err error)

	// 启用设备
	Enable(ctx context.Context, ids []int64) (err error)
	// 禁用设备
	Disable(ctx context.Context, ids []int64) (err error)
	// 设备上线
	Online(ctx context.Context, in *v1.DeviceOnlineReq) (err error)
	// 设备下线
	Offline(ctx context.Context, in *v1.DeviceOfflineReq) (err error)
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
