package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IDeviceConfig interface {
	Add(ctx context.Context, in *v1.AddDeviceConfigReq) (err error)

	Delete(ctx context.Context, id int64) (err error)

	Edit(ctx context.Context, in *v1.EditDeviceConfigReq) (err error)

	Get(ctx context.Context, id int64) (out *model.DeviceConfig, err error)
	List(ctx context.Context, in *v1.ListDeviceConfigReq) (out []*model.DeviceConfig, err error)
}

var localDeviceConfig IDeviceConfig

func DeviceConfig() IDeviceConfig {
	if localDeviceConfig == nil {
		panic("implement not found for interface IDeviceConfig, forgot register?")
	}
	return localDeviceConfig
}

func RegisterDeviceConfig(i IDeviceConfig) {
	localDeviceConfig = i
}
