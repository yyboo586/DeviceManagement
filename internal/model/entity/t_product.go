package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TProduct 产品实体
type TProduct struct {
	ID         int64       `orm:"id"`          // 主键
	OrgID      string      `orm:"org_id"`      // 组织ID
	CategoryID int64       `orm:"category_id"` // 产品分类ID
	Name       string      `orm:"name"`        // 产品名称
	Desc       string      `orm:"desc"`        // 产品描述
	CreatedAt  *gtime.Time `orm:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at"`  // 修改时间
}
