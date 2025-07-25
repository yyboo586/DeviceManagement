package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type deviceLogController struct{}

var DeviceLogController = &deviceLogController{}

func (c *deviceLogController) List(ctx context.Context, req *v1.DeviceLogListReq) (res *v1.DeviceLogListRes, err error) {
	logs, pageRes, err := service.DeviceLog().List(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.DeviceLogListRes{
		PageRes: pageRes,
	}
	for _, v := range logs {
		res.List = append(res.List, c.format(v))
	}

	return
}

func (c *deviceLogController) format(in *model.DeviceLog) (out *v1.DeviceLog) {
	out = &v1.DeviceLog{
		ID:         in.ID,
		OrgID:      in.OrgID,
		DeviceID:   in.DeviceID,
		DeviceName: in.DeviceName,
		DeviceKey:  in.DeviceKey,
		Type:       model.GetDeviceLogType(in.Type),
		Timestamp:  gtime.New(time.Unix(in.Timestamp, 0)).String(),
		CreatedAt:  gtime.New(time.Unix(in.CreatedAt, 0)).String(),
	}
	return
}
