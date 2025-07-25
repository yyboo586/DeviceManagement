package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

type deviceConfigController struct{}

var DeviceConfigController = &deviceConfigController{}

func (c *deviceConfigController) Add(ctx context.Context, req *v1.AddDeviceConfigReq) (res *v1.AddDeviceConfigRes, err error) {
	err = service.DeviceConfig().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (c *deviceConfigController) Delete(ctx context.Context, req *v1.DeleteDeviceConfigReq) (res *v1.DeleteDeviceConfigRes, err error) {
	err = service.DeviceConfig().Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return
}

func (c *deviceConfigController) Edit(ctx context.Context, req *v1.EditDeviceConfigReq) (res *v1.EditDeviceConfigRes, err error) {
	err = service.DeviceConfig().Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (c *deviceConfigController) Get(ctx context.Context, req *v1.GetDeviceConfigReq) (res *v1.GetDeviceConfigRes, err error) {
	out, err := service.DeviceConfig().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetDeviceConfigRes{
		DeviceConfig: c.formmat(out),
	}
	return
}

func (c *deviceConfigController) List(ctx context.Context, req *v1.ListDeviceConfigReq) (res *v1.ListDeviceConfigRes, err error) {
	out, err := service.DeviceConfig().List(ctx, req)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.DeviceConfig, 0)
	for _, item := range out {
		list = append(list, c.formmat(item))
	}

	res = &v1.ListDeviceConfigRes{
		List: list,
	}
	return
}

func (c *deviceConfigController) formmat(in *model.DeviceConfig) (out *v1.DeviceConfig) {
	out = &v1.DeviceConfig{
		ID:        in.ID,
		OrgID:     in.OrgID,
		Type:      model.GetDeviceConfigTypeText(in.Type),
		Key:       in.Key,
		Value:     in.Value,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
