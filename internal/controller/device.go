package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/logics"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

type deviceController struct{}

var DeviceController = deviceController{}

func (c *deviceController) Add(ctx context.Context, req *v1.AddDeviceReq) (res *v1.AddDeviceRes, err error) {
	in := &model.Device{
		Name:        req.Name,
		DeviceKey:   req.DeviceKey,
		OrgID:       req.OrgID,
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
		Device: device,
	}

	return res, nil
}

func (c *deviceController) List(ctx context.Context, req *v1.ListDeviceReq) (res *v1.ListDeviceRes, err error) {
	devices, pageRes, err := service.Device().List(ctx, req.OrgID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ListDeviceRes{}
	res.List = append(res.List, devices...)
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

func (c *deviceController) Online(ctx context.Context, req *v1.DeviceOnlineReq) (res *v1.DeviceOnlineRes, err error) {
	err = service.Device().Online(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.DeviceOnlineRes{}, nil
}

func (c *deviceController) Offline(ctx context.Context, req *v1.DeviceOfflineReq) (res *v1.DeviceOfflineRes, err error) {
	err = service.Device().Offline(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.DeviceOfflineRes{}, nil
}

func (c *deviceController) LogList(ctx context.Context, req *v1.DeviceOnlineLogListReq) (res *v1.DeviceOnlineLogListRes, err error) {
	deviceOnlineLog := logics.NewDeviceOnlineLog()

	listReq := &model.DeviceOnlineLogListReq{
		OrgID:     req.OrgID,
		DeviceID:  req.DeviceID,
		EventType: req.EventType,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageReq: &model.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}

	logs, pageRes, err := deviceOnlineLog.List(ctx, listReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DeviceOnlineLogListRes{
		List:    logs,
		PageRes: pageRes,
	}

	return
}
