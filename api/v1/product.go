package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddProductReq struct {
	g.Meta     `path:"/products" tags:"分类管理/产品管理" method:"post" summary:"新增"`
	CategoryID int64  `json:"category_id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
	Desc       string `json:"desc" dc:"产品描述"`
}

type AddProductRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"产品ID"`
}

type DeleteProductReq struct {
	g.Meta `path:"/products" tags:"分类管理/产品管理" method:"delete" summary:"删除"`
	IDs    []int64 `json:"ids" v:"required#产品ID不能为空" dc:"产品ID"`
}

type DeleteProductRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateProductReq struct {
	g.Meta     `path:"/products/{id}" tags:"分类管理/产品管理" method:"put" summary:"编辑"`
	ID         int64  `p:"id" v:"required#产品ID不能为空" dc:"产品ID"`
	CategoryID int64  `json:"category_id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
	Desc       string `json:"desc" dc:"产品描述"`
}

type UpdateProductRes struct {
	g.Meta `mime:"application/json"`
}

type GetProductReq struct {
	g.Meta `path:"/products/{id}" tags:"分类管理/产品管理" method:"get" summary:"详情"`
	ID     int64 `p:"id" v:"required#产品ID不能为空" dc:"产品ID"`
}

type GetProductRes struct {
	g.Meta `mime:"application/json"`
	*Product
}

type ListProductReq struct {
	g.Meta     `path:"/products" tags:"分类管理/产品管理" method:"get" summary:"列表"`
	CategoryID int64 `json:"category_id" dc:"产品分类ID"`
	model.PageReq
}

type ListProductRes struct {
	g.Meta `mime:"application/json"`
	List   []*Product `json:"list" dc:"产品列表"`
	*model.PageRes
}

type Product struct {
	ID         int64       `json:"id" dc:"产品ID"`            // 主键
	CategoryID int64       `json:"category_id" dc:"产品分类ID"` // 产品分类ID
	Name       string      `json:"name" dc:"产品名称"`          // 产品名称
	Desc       string      `json:"desc" dc:"产品描述"`          // 产品描述
	CreatedAt  *gtime.Time `json:"created_at" dc:"创建时间"`    // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at" dc:"修改时间"`    // 修改时间
}
