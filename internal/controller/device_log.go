package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/service"
	"context"
)

type deviceLogController struct{}

var DeviceLog = deviceLogController{}

/*
func (c *deviceLogController) Add(ctx context.Context, req *v1.AddDeviceLogReq) (res *v1.AddDeviceLogRes, err error) {
	err = service.DeviceLog().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.Type == int(model.DeviceLogTypeAlarm) {
		err = service.Device().Alarm(ctx, &v1.DeviceAlarmReq{
			OrgID:    req.OrgID,
			DeviceID: req.DeviceID,
			Content:  req.Content.Message,
		})
	}
	return
}
*/

func (c *deviceLogController) List(ctx context.Context, req *v1.DeviceLogListReq) (res *v1.DeviceLogListRes, err error) {
	logs, pageRes, err := service.DeviceLog().List(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.DeviceLogListRes{
		List:    logs,
		PageRes: pageRes,
	}

	return
}
