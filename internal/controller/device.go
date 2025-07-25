package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

type deviceController struct{}

var DeviceController = &deviceController{}

func (c *deviceController) Add(ctx context.Context, req *v1.AddDeviceReq) (res *v1.AddDeviceRes, err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	in := &model.Device{
		ProductID:   req.ProductID,
		OrgID:       operatorInfo.OrgID,
		CreatorID:   operatorInfo.UserID,
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

func (c *deviceController) EditDeviceStatus(ctx context.Context, req *v1.EditDeviceStatusReq) (res *v1.EditDeviceStatusRes, err error) {
	err = service.Device().EditDeviceStatus(ctx, req.IDs, req.Enabled)
	if err != nil {
		return nil, err
	}

	return &v1.EditDeviceStatusRes{}, nil
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
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	devices, pageRes, err := service.Device().List(ctx, operatorInfo.OrgID, &req.PageReq)
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

func (c *deviceController) format(in *model.Device) (out *v1.Device) {
	return &v1.Device{
		ID:              in.ID,
		Name:            in.Name,
		DeviceKey:       in.DeviceKey,
		OrgID:           in.OrgID,
		CreatorID:       in.CreatorID,
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
