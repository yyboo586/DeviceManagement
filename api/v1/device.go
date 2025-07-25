package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AddDeviceReq struct {
	g.Meta `path:"/devices" tags:"设备管理" method:"post" summary:"添加设备"`
	model.Author
	OrgID       string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
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
	*model.Device
}

type ListDeviceReq struct {
	g.Meta `path:"/devices" tags:"设备管理" method:"get" summary:"设备列表(根据组织ID)"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	model.PageReq
}

type ListDeviceRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.Device `json:"list" dc:"设备列表"`
	*model.PageRes
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

type DeviceOnlineReq struct {
	g.Meta `path:"/devices/status/online" tags:"设备管理/设备状态" method:"put" summary:"设备上线"`
	model.Author
	OrgID   string                `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Devices []*DeviceOnlineStatus `json:"devices" v:"required#设备列表不能为空" dc:"设备列表"`
}

type DeviceOnlineRes struct {
	g.Meta `mime:"application/json"`
}

type DeviceOfflineReq struct {
	g.Meta `path:"/devices/status/offline" tags:"设备管理/设备状态" method:"put" summary:"设备下线"`
	model.Author
	OrgID   string                `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Devices []*DeviceOnlineStatus `json:"devices" v:"required#设备列表不能为空" dc:"设备列表"`
}

type DeviceOfflineRes struct {
	g.Meta `mime:"application/json"`
}

type DeviceOnlineStatus struct {
	DeviceID  int64  `json:"device_id" v:"required#设备ID不能为空" dc:"设备ID"`
	DeviceKey string `json:"device_key" v:"required#设备唯一标识不能为空" dc:"设备唯一标识"`
}
