package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddDeviceReq struct {
	g.Meta `path:"/devices" tags:"设备管理" method:"post" summary:"添加设备"`
	model.Author
	OrgID       string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	ProductID   int64  `json:"product_id" v:"required#产品ID不能为空" dc:"产品ID"`
	Name        string `json:"name" v:"required#设备名称不能为空" dc:"设备名称"`
	DeviceKey   string `json:"device_key" v:"required#设备唯一标识不能为空" dc:"设备唯一标识"`
	Location    string `json:"location" dc:"设备位置"`
	Description string `json:"description" dc:"设备描述"`
}

type AddDeviceRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"设备ID"`
}

type DeleteDeviceReq struct {
	g.Meta `path:"/devices" tags:"设备管理" method:"delete" summary:"删除设备"`
	model.Author
	IDs []int64 `json:"ids" v:"required#设备IDs不能为空" dc:"设备IDs"`
}

type DeleteDeviceRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateDeviceReq struct {
	g.Meta `path:"/devices/{id}" tags:"设备管理" method:"put" summary:"更新设备"`
	model.Author
	ID          int64  `p:"id" v:"required#设备ID不能为空" dc:"设备ID"`
	Name        string `json:"name" dc:"设备名称"`
	Location    string `json:"location" dc:"设备位置"`
	Description string `json:"description" dc:"设备描述"`
}

type UpdateDeviceRes struct {
	g.Meta `mime:"application/json"`
}

type GetDeviceReq struct {
	g.Meta `path:"/devices/{id}" tags:"设备管理" method:"get" summary:"设备详情"`
	model.Author
	ID int64 `p:"id" v:"required#设备ID不能为空" dc:"设备ID"`
}

type GetDeviceRes struct {
	g.Meta `mime:"application/json"`
	*Device
}

type ListDeviceReq struct {
	g.Meta `path:"/devices" tags:"设备管理" method:"get" summary:"设备列表(根据组织ID)"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	model.PageReq
}

type ListDeviceRes struct {
	g.Meta `mime:"application/json"`
	List   []*Device `json:"list" dc:"设备列表"`
	*model.PageRes
}

type Device struct {
	ID              int64       `json:"id" dc:"设备ID"`
	Name            string      `json:"name" dc:"设备名称"`
	DeviceKey       string      `json:"device_key" dc:"设备唯一标识"`
	OrgID           string      `json:"org_id" dc:"组织ID"`
	Enabled         bool        `json:"enabled" dc:"设备状态"`
	OnlineStatus    string      `json:"online_status" dc:"设备在线状态"`
	Location        string      `json:"location" dc:"设备位置"`
	Description     string      `json:"description" dc:"设备描述"`
	LastOnlineTime  *gtime.Time `json:"last_online_time" dc:"最后在线时间"`
	LastOfflineTime *gtime.Time `json:"last_offline_time" dc:"最后离线时间"`
	CreatedAt       *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updated_at" dc:"更新时间"`
}

type DeviceEnableReq struct {
	g.Meta `path:"/devices/status/enable" tags:"设备管理/设备状态" method:"put" summary:"启用设备"`
	model.Author
	IDs []int64 `json:"ids" v:"required#设备IDs不能为空" dc:"设备IDs"`
}

type DeviceEnableRes struct {
	g.Meta `mime:"application/json"`
}

type DeviceDisableReq struct {
	g.Meta `path:"/devices/status/disable" tags:"设备管理/设备状态" method:"put" summary:"禁用设备"`
	model.Author
	IDs []int64 `json:"ids" v:"required#设备IDs不能为空" dc:"设备IDs"`
}

type DeviceDisableRes struct {
	g.Meta `mime:"application/json"`
}

/*
type DeviceAlarmReq struct {
	g.Meta `path:"/devices/alarm" tags:"设备管理" method:"post" summary:"设备告警"`
	model.Author
	OrgID    string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	DeviceID int64  `json:"device_id" v:"required#设备ID不能为空" dc:"设备ID"`
	Content  string `json:"content" v:"required#告警内容不能为空" dc:"告警内容"`
}

type DeviceAlarmRes struct {
	g.Meta `mime:"application/json"`
}
*/
