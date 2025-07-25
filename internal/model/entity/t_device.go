package entity

import "github.com/gogf/gf/v2/os/gtime"

type TDevice struct {
	ID              int64       `orm:"id"`
	ProductID       int64       `orm:"product_id"`
	OrgID           string      `orm:"org_id"`
	CreatorID       string      `orm:"creator_id"`
	DeviceKey       string      `orm:"device_key"`
	Name            string      `orm:"name"`
	Enabled         int         `orm:"enabled"`
	OnlineStatus    int         `orm:"online_status"`
	Location        string      `orm:"location"`
	Description     string      `orm:"description"`
	LastOnlineTime  *gtime.Time `orm:"last_online_time"`
	LastOfflineTime *gtime.Time `orm:"last_offline_time"`
	CreatedAt       *gtime.Time `orm:"created_at"`
	UpdatedAt       *gtime.Time `orm:"updated_at"`
}
