package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TProductCategory 产品分类实体
type TProductCategory struct {
	ID        int64       `orm:"id"`         // 主键
	PID       int64       `orm:"pid"`        // 父级分类ID
	OrgID     string      `orm:"org_id"`     // 组织ID
	Name      string      `orm:"name"`       // 产品分类名称
	CreatedAt *gtime.Time `orm:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at"` // 修改时间
}
