package service

import (
	v1 "DeviceManagement/api/v1"
	"context"
)

type IDevicePermissionService interface {
	BindDevicePermission(ctx context.Context, req *v1.BindDevicePermissionReq) error
}

var localDevicePermissionService IDevicePermissionService

func DevicePermission() IDevicePermissionService {
	if localDevicePermissionService == nil {
		panic("implement not found for interface IDevicePermissionService, forgot register?")
	}
	return localDevicePermissionService
}

func RegisterDevicePermission(i IDevicePermissionService) {
	localDevicePermissionService = i
}
