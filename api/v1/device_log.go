package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DeviceLogListReq struct {
	g.Meta `path:"/devices/logs" tags:"设备管理/日志管理" method:"get" summary:"列表"`
	model.Author
	DeviceID  int64       `json:"device_id" dc:"设备ID"`
	Type      int         `json:"type" dc:"日志类型(1:上线 2:下线 3:报警)"`
	StartTime *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" dc:"结束时间"`
	model.PageReq
}

type DeviceLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*DeviceLog `json:"list" dc:"日志列表"`
	*model.PageRes
}

type DeviceLog struct {
	ID         int64  `json:"id" dc:"日志ID"`
	OrgID      string `json:"org_id" dc:"组织ID"`
	DeviceID   int64  `json:"device_id" dc:"设备ID"`
	DeviceName string `json:"device_name" dc:"设备名称"`
	DeviceKey  string `json:"device_key" dc:"设备唯一标识"`
	Type       string `json:"type" dc:"日志类型"`
	Timestamp  string `json:"timestamp" dc:"时间戳"`
	CreatedAt  string `json:"created_at" dc:"创建时间"`
}
