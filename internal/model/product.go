package model

import "github.com/gogf/gf/v2/os/gtime"

type Product struct {
	ID         int64       `json:"id" dc:"产品ID"`            // 主键
	OrgID      string      `json:"org_id" dc:"组织ID"`        // 组织ID
	CategoryID int64       `json:"category_id" dc:"产品分类ID"` // 产品分类ID
	Name       string      `json:"name" dc:"产品名称"`          // 产品名称
	Desc       string      `json:"desc" dc:"产品描述"`          // 产品描述
	CreatedAt  *gtime.Time `json:"created_at" dc:"创建时间"`    // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at" dc:"修改时间"`    // 修改时间
}
