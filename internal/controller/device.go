package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

type deviceController struct{}

var Device = deviceController{}

func (c *deviceController) Add(ctx context.Context, req *v1.AddDeviceReq) (res *v1.AddDeviceRes, err error) {
	in := &model.Device{
		ProductID:   req.ProductID,
		OrgID:       req.OrgID,
		DeviceKey:   req.DeviceKey,
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
	}

	id, err := service.Device().Add(ctx, in)
	if err != nil {
		return nil, err
	}

	return &v1.AddDeviceRes{ID: id}, nil
}

func (c *deviceController) Delete(ctx context.Context, req *v1.DeleteDeviceReq) (res *v1.DeleteDeviceRes, err error) {
	err = service.Device().Delete(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDeviceRes{}, nil
}

func (c *deviceController) Update(ctx context.Context, req *v1.UpdateDeviceReq) (res *v1.UpdateDeviceRes, err error) {
	in := &model.Device{
		ID:          req.ID,
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
	}

	err = service.Device().Update(ctx, in)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDeviceRes{}, nil
}

func (c *deviceController) Get(ctx context.Context, req *v1.GetDeviceReq) (res *v1.GetDeviceRes, err error) {
	device, err := service.Device().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetDeviceRes{
		Device: c.format(device),
	}

	return res, nil
}

func (c *deviceController) List(ctx context.Context, req *v1.ListDeviceReq) (res *v1.ListDeviceRes, err error) {
	devices, pageRes, err := service.Device().List(ctx, req.OrgID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListDeviceRes{}
	for _, device := range devices {
		res.List = append(res.List, c.format(device))
	}
	res.PageRes = pageRes
	return
}

func (c *deviceController) Enable(ctx context.Context, req *v1.DeviceEnableReq) (res *v1.DeviceEnableRes, err error) {
	err = service.Device().Enable(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	return &v1.DeviceEnableRes{}, nil
}

func (c *deviceController) Disable(ctx context.Context, req *v1.DeviceDisableReq) (res *v1.DeviceDisableRes, err error) {
	err = service.Device().Disable(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	return &v1.DeviceDisableRes{}, nil
}

/*
func (c *deviceController) Alarm(ctx context.Context, req *v1.DeviceAlarmReq) (res *v1.DeviceAlarmRes, err error) {
	err = service.Device().Alarm(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.DeviceAlarmRes{}, nil
}
*/

func (c *deviceController) format(in *model.Device) (out *v1.Device) {
	return &v1.Device{
		ID:              in.ID,
		Name:            in.Name,
		DeviceKey:       in.DeviceKey,
		OrgID:           in.OrgID,
		Enabled:         in.Enabled == model.DeviceStatusEnabled,
		OnlineStatus:    model.GetDeviceOnlineStatusText(in.OnlineStatus),
		Location:        in.Location,
		Description:     in.Description,
		LastOnlineTime:  in.LastOnlineTime,
		LastOfflineTime: in.LastOfflineTime,
		CreatedAt:       in.CreatedAt,
		UpdatedAt:       in.UpdatedAt,
	}
}
