package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// DeviceOnlineLogListReq 设备上下线日志列表请求
type DeviceOnlineLogListReq struct {
	g.Meta `path:"/devices/logs/online_status" tags:"设备管理/日志管理" method:"get" summary:"获取设备上下线日志列表"`
	model.Author
	OrgID     string `json:"org_id" dc:"组织ID" v:"required"`
	DeviceID  int64  `json:"device_id" dc:"设备ID"`
	EventType int    `json:"event_type" dc:"事件类型(1:上线 2:下线)"`
	StartTime string `json:"start_time" dc:"开始时间"`
	EndTime   string `json:"end_time" dc:"结束时间"`
	model.PageReq
}

// DeviceOnlineLogListRes 设备上下线日志列表响应
type DeviceOnlineLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.DeviceOnlineLog `json:"list" dc:"日志列表"`
	*model.PageRes
}

// DeviceOnlineDurationReq 设备在线时长请求
type DeviceOnlineDurationReq struct {
	g.Meta `path:"/devices/logs/online_duration" tags:"设备管理/日志管理" method:"get" summary:"获取设备在线时长"`
	model.Author
	DeviceID  int64  `json:"device_id" dc:"设备ID" v:"required"`
	StartTime string `json:"start_time" dc:"开始时间"`
	EndTime   string `json:"end_time" dc:"结束时间"`
}

// DeviceOnlineDurationRes 设备在线时长响应
type DeviceOnlineDurationRes struct {
	g.Meta    `mime:"application/json"`
	DeviceID  int64  `json:"device_id" dc:"设备ID"`
	Duration  int64  `json:"duration" dc:"在线时长(秒)"`
	StartTime string `json:"start_time" dc:"开始时间"`
	EndTime   string `json:"end_time" dc:"结束时间"`
}
