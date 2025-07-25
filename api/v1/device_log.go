package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*
	type AddDeviceLogReq struct {
		g.Meta `path:"/devices/logs" tags:"设备管理/日志管理" method:"post" summary:"添加设备日志"`
		model.Author
		OrgID     string                  `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
		DeviceID  int64                   `json:"device_id" v:"required#设备ID不能为空" dc:"设备ID"`
		DeviceKey string                  `json:"device_key" v:"required#设备唯一标识不能为空" dc:"设备唯一标识"`
		Type      int                     `json:"type" v:"required#日志类型不能为空" dc:"日志类型(1:上线 2:下线 3:报警)"`
		Content   *model.DeviceLogContent `json:"content" dc:"日志内容"`
		CreatedAt *gtime.Time             `json:"created_at" v:"required#创建时间不能为空" dc:"创建时间"`
	}

	type AddDeviceLogRes struct {
		g.Meta `mime:"application/json"`
	}
*/
type DeviceLogListReq struct {
	g.Meta `path:"/devices/logs" tags:"设备管理/日志管理" method:"get" summary:"获取设备日志列表"`
	model.Author
	OrgID     string      `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	DeviceID  int64       `json:"device_id" dc:"设备ID"`
	Type      int         `json:"type" dc:"日志类型(1:上线 2:下线 3:报警)"`
	StartTime *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" dc:"结束时间"`
	model.PageReq
}

type DeviceLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.DeviceLog `json:"list" dc:"日志列表"`
	*model.PageRes
}
