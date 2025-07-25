package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type devicePermissionController struct {
}

var DevicePermissionController = &devicePermissionController{}

func (c *devicePermissionController) BindDevicePermission(ctx context.Context, req *v1.BindDevicePermissionReq) (res *v1.BindDevicePermissionRes, err error) {
	err = service.DevicePermission().BindDevicePermission(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}
