package model

// 设备日志类型
type DeviceLogType int

const (
	_                    DeviceLogType = iota
	DeviceLogTypeOnline                // 上线
	DeviceLogTypeOffline               // 下线
	DeviceLogTypeAlarm                 // 报警
)

func GetDeviceLogType(in DeviceLogType) string {
	switch in {
	case DeviceLogTypeOnline:
		return "上线"
	case DeviceLogTypeOffline:
		return "下线"
	case DeviceLogTypeAlarm:
		return "报警"
	default:
		return "未知"
	}
}

// DeviceLog 设备日志
type DeviceLog struct {
	ID         int64             `json:"id" dc:"日志ID"`
	OrgID      string            `json:"org_id" dc:"组织ID"`
	DeviceID   int64             `json:"device_id" dc:"设备ID"`
	DeviceName string            `json:"device_name" dc:"设备名称"`
	DeviceKey  string            `json:"device_key" dc:"设备唯一标识"`
	Type       DeviceLogType     `json:"type" dc:"日志类型"`
	Content    *DeviceLogContent `json:"content" dc:"日志内容"`
	Timestamp  int64             `json:"timestamp" dc:"时间戳"`
	CreatedAt  int64             `json:"created_at" dc:"创建时间"`
}

type DeviceLogContent struct {
	Message string                 `json:"message" dc:"消息"`
	Details map[string]interface{} `json:"details" dc:"详情"`
}
