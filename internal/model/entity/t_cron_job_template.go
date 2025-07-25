package entity

import "github.com/gogf/gf/v2/os/gtime"

type TCronJobTemplate struct {
	ID         string      `orm:"id"`
	Name       string      `orm:"name"`
	InvokeType string      `orm:"invoke_type"`
	Config     string      `orm:"config"`
	CreatedAt  *gtime.Time `orm:"created_at"`
	UpdatedAt  *gtime.Time `orm:"updated_at"`
}
