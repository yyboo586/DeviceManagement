package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/logics"
	"DeviceManagement/internal/model"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type cDeviceOnlineLog struct{}

var DeviceOnlineLog = cDeviceOnlineLog{}

// List 获取设备上下线日志列表
func (c *cDeviceOnlineLog) List(ctx context.Context, req *v1.DeviceOnlineLogListReq) (res *v1.DeviceOnlineLogListRes, err error) {
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

// GetOnlineDuration 获取设备在线时长
func (c *cDeviceOnlineLog) GetOnlineDuration(ctx context.Context, req *v1.DeviceOnlineDurationReq) (res *v1.DeviceOnlineDurationRes, err error) {
	deviceOnlineLog := logics.NewDeviceOnlineLog()

	var startTime, endTime *gtime.Time
	if req.StartTime != "" {
		startTime = gtime.NewFromStr(req.StartTime)
	}
	if req.EndTime != "" {
		endTime = gtime.NewFromStr(req.EndTime)
	}

	duration, err := deviceOnlineLog.GetDeviceOnlineDuration(ctx, req.DeviceID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	res = &v1.DeviceOnlineDurationRes{
		DeviceID:  req.DeviceID,
		Duration:  duration,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	return
}
