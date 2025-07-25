package model

import "github.com/gogf/gf/v2/os/gtime"

// 设备上下线事件类型
type DeviceOnlineEventType int

const (
	_                        DeviceOnlineEventType = iota
	DeviceOnlineEventOnline                        // 上线
	DeviceOnlineEventOffline                       // 下线
)

// DeviceOnlineLog 设备上下线日志
type DeviceOnlineLog struct {
	ID           int64                 `json:"id" dc:"日志ID"`
	DeviceID     int64                 `json:"device_id" dc:"设备ID"`
	DeviceKey    string                `json:"device_key" dc:"设备唯一标识"`
	OrgID        string                `json:"org_id" dc:"组织ID"`
	EventType    DeviceOnlineEventType `json:"event_type" dc:"事件类型"`
	OnlineStatus DeviceOnlineStatus    `json:"online_status" dc:"设备在线状态"`
	IPAddress    string                `json:"ip_address" dc:"设备IP地址"`
	ClientID     string                `json:"client_id" dc:"客户端ID"`
	Reason       string                `json:"reason" dc:"上下线原因"`
	Duration     int64                 `json:"duration" dc:"在线时长(秒)"`
	CreatedAt    *gtime.Time           `json:"created_at" dc:"创建时间"`
}

// DeviceOnlineLogReq 设备上下线日志请求
type DeviceOnlineLogReq struct {
	DeviceID     int64                 `json:"device_id" dc:"设备ID"`
	DeviceKey    string                `json:"device_key" dc:"设备唯一标识"`
	OrgID        string                `json:"org_id" dc:"组织ID"`
	EventType    DeviceOnlineEventType `json:"event_type" dc:"事件类型"`
	OnlineStatus DeviceOnlineStatus    `json:"online_status" dc:"设备在线状态"`
	IPAddress    string                `json:"ip_address" dc:"设备IP地址"`
	ClientID     string                `json:"client_id" dc:"客户端ID"`
	Reason       string                `json:"reason" dc:"上下线原因"`
	Duration     int64                 `json:"duration" dc:"在线时长(秒)"`
}

// DeviceOnlineLogListReq 设备上下线日志列表请求
type DeviceOnlineLogListReq struct {
	OrgID     string `json:"org_id" dc:"组织ID"`
	DeviceID  int64  `json:"device_id" dc:"设备ID"`
	EventType int    `json:"event_type" dc:"事件类型"`
	StartTime string `json:"start_time" dc:"开始时间"`
	EndTime   string `json:"end_time" dc:"结束时间"`
	*PageReq
}
