package entity

type TDeviceLog struct {
	ID        int64  `orm:"id"`
	OrgID     string `orm:"org_id"`
	DeviceID  int64  `orm:"device_id"`
	DeviceKey string `orm:"device_key"`
	Type      int    `orm:"type"`
	Content   string `orm:"content"`
	CreatedAt int64  `orm:"created_at"`
}
