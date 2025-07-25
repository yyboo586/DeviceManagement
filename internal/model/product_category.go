package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ProductCategory struct {
	ID        int64       `json:"id" dc:"分类ID"`         // 主键
	PID       int64       `json:"pid" dc:"父级分类ID"`      // 父级分类ID
	OrgID     string      `json:"org_id" dc:"组织ID"`     // 组织ID
	Name      string      `json:"name" dc:"名称"`         // 产品分类名称
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" dc:"修改时间"` // 修改时间
}

type ProductCategoryTree struct {
	*ProductCategory
	Children []*ProductCategoryTree `json:"children" dc:"子分类"`
}

type ProductCategoryTreeList []*ProductCategoryTree
