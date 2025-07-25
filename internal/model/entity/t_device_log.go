package entity

type TDeviceLog struct {
	ID         int64  `orm:"id"`
	OrgID      string `orm:"org_id"`
	DeviceID   int64  `orm:"device_id"`
	DeviceName string `orm:"device_name"`
	DeviceKey  string `orm:"device_key"`
	Type       int    `orm:"type"`
	Content    string `orm:"content"`
	Timestamp  int64  `orm:"timestamp"`
	CreatedAt  int64  `orm:"created_at"`
}
