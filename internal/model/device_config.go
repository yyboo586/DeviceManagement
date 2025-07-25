package model

import "github.com/gogf/gf/v2/os/gtime"

type DeviceConfigType int

const (
	DeviceConfigTypeUnknown DeviceConfigType = iota
	DeviceConfigTypeAlarm                    // 告警
)

func GetDeviceConfigType(typ string) DeviceConfigType {
	switch typ {
	case "alarm":
		return DeviceConfigTypeAlarm
	default:
		return DeviceConfigTypeUnknown
	}
}

func GetDeviceConfigTypeText(typ DeviceConfigType) string {
	switch typ {
	case DeviceConfigTypeAlarm:
		return "alarm"
	default:
		return "unknown"
	}
}

type DeviceConfig struct {
	ID        int64            `json:"id" dc:"配置ID"`
	OrgID     string           `json:"org_id" dc:"组织ID"`
	Type      DeviceConfigType `json:"type" dc:"配置类型"`
	Key       string           `json:"key" dc:"配置键"`
	Value     string           `json:"value" dc:"配置值"`
	CreatedAt *gtime.Time      `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time      `json:"updated_at" dc:"更新时间"`
}
