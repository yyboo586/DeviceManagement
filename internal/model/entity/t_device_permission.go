package entity

type DevicePermission struct {
	ID        int64  `orm:"id"`
	OrgID     string `orm:"org_id"`
	UserID    string `orm:"user_id"`
	DeviceID  int64  `orm:"device_id"`
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}
