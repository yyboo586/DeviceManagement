package entity

import "github.com/gogf/gf/v2/os/gtime"

type TDeviceConfig struct {
	ID        int64       `orm:"id" dc:"配置ID"`
	OrgID     string      `orm:"org_id" dc:"组织ID"`
	Type      int         `orm:"type" dc:"配置类型"`
	Key       string      `orm:"key" dc:"配置键"`
	Value     string      `orm:"value" dc:"配置值"`
	CreatedAt *gtime.Time `orm:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" dc:"更新时间"`
}
