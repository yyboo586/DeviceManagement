package entity

import "github.com/gogf/gf/v2/os/gtime"

type TDeviceOnlineLog struct {
	ID           int64       `orm:"id"`
	DeviceID     int64       `orm:"device_id"`
	DeviceKey    string      `orm:"device_key"`
	OrgID        string      `orm:"org_id"`
	EventType    int         `orm:"event_type"`
	OnlineStatus int         `orm:"online_status"`
	IPAddress    string      `orm:"ip_address"`
	ClientID     string      `orm:"client_id"`
	Reason       string      `orm:"reason"`
	Duration     int64       `orm:"duration"`
	CreatedAt    *gtime.Time `orm:"created_at"`
}
