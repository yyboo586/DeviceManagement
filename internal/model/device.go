package model

import "github.com/gogf/gf/v2/os/gtime"

type DeviceStatus int

const (
	_                    DeviceStatus = iota
	DeviceStatusEnabled               // 启用
	DeviceStatusDisabled              // 禁用
)

type DeviceOnlineStatus int

const (
	_                         DeviceOnlineStatus = iota
	DeviceOnlineStatusOnline                     // 在线
	DeviceOnlineStatusOffline                    // 离线
)

func GetDeviceOnlineStatusText(status DeviceOnlineStatus) string {
	switch status {
	case DeviceOnlineStatusOnline:
		return "在线"
	case DeviceOnlineStatusOffline:
		return "离线"
	default:
		return "未知"
	}
}

type Device struct {
	ID              int64              `json:"id" dc:"设备ID"`
	ProductID       int64              `json:"product_id" dc:"产品ID"`
	OrgID           string             `json:"org_id" dc:"组织ID"`
	CreatorID       string             `json:"creator_id" dc:"创建者ID"`
	DeviceKey       string             `json:"device_key" dc:"设备唯一标识"`
	Name            string             `json:"name" dc:"设备名称"`
	Enabled         DeviceStatus       `json:"enabled" dc:"设备状态"`
	OnlineStatus    DeviceOnlineStatus `json:"online_status" dc:"设备在线状态"`
	Location        string             `json:"location" dc:"设备位置"`
	Description     string             `json:"description" dc:"设备描述"`
	LastOnlineTime  *gtime.Time        `json:"last_online_time" dc:"最后在线时间"`
	LastOfflineTime *gtime.Time        `json:"last_offline_time" dc:"最后离线时间"`
	CreatedAt       *gtime.Time        `json:"created_at" dc:"创建时间"`
	UpdatedAt       *gtime.Time        `json:"updated_at" dc:"更新时间"`
}
