package model

import "github.com/gogf/gf/v2/os/gtime"

// 设备日志类型
type DeviceLogType int

const (
	_                    DeviceLogType = iota
	DeviceLogTypeOnline                // 上线
	DeviceLogTypeOffline               // 下线
	DeviceLogTypeAlarm                 // 报警
)

// DeviceLog 设备日志
type DeviceLog struct {
	ID        int64             `json:"id" dc:"日志ID"`
	OrgID     string            `json:"org_id" dc:"组织ID"`
	DeviceID  int64             `json:"device_id" dc:"设备ID"`
	DeviceKey string            `json:"device_key" dc:"设备唯一标识"`
	Type      DeviceLogType     `json:"type" dc:"日志类型"`
	Content   *DeviceLogContent `json:"content" dc:"日志内容"`
	CreatedAt *gtime.Time       `json:"created_at" dc:"创建时间"`
}

type DeviceLogContent struct {
	Message string                 `json:"message" dc:"消息"`
	Details map[string]interface{} `json:"details" dc:"详情"`
}
