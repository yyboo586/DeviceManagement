package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ProductCategory struct {
	ID        int64       `json:"id" dc:"分类ID"`         // 主键
	OrgID     string      `json:"org_id" dc:"组织ID"`     // 组织ID
	Name      string      `json:"name" dc:"名称"`         // 产品分类名称
	Desc      string      `json:"desc" dc:"描述"`         // 产品分类描述
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" dc:"修改时间"` // 修改时间
}
