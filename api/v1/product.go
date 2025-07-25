package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AddProductReq struct {
	g.Meta     `path:"/products" tags:"产品管理" method:"post" summary:"新增"`
	CategoryID int64  `json:"category_id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	OrgID      string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
}

type AddProductRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"产品ID"`
}

type DeleteProductReq struct {
	g.Meta `path:"/products/{id}" tags:"产品管理" method:"delete" summary:"删除"`
	ID     int64 `p:"id" v:"required#产品ID不能为空" dc:"产品ID"`
}

type DeleteProductRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateProductReq struct {
	g.Meta     `path:"/products/{id}" tags:"产品管理" method:"put" summary:"修改(全量更新)"`
	ID         int64  `p:"id" v:"required#产品ID不能为空" dc:"产品ID"`
	CategoryID int64  `json:"category_id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
}

type UpdateProductRes struct {
	g.Meta `mime:"application/json"`
}

type GetProductReq struct {
	g.Meta `path:"/products/{id}" tags:"产品管理" method:"get" summary:"根据ID查询"`
	ID     int64 `p:"id" v:"required#产品ID不能为空" dc:"产品ID"`
}

type GetProductRes struct {
	g.Meta `mime:"application/json"`
	*model.Product
}

type ListProductReq struct {
	g.Meta     `path:"/products" tags:"产品管理" method:"get" summary:"查询产品列表"`
	OrgID      string `json:"org_id" v:"required#组织ID不能为空"`
	CategoryID int64  `json:"category_id" dc:"产品分类ID"`
}

type ListProductRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.Product `json:"list" dc:"产品列表"`
	Total  int              `json:"total" dc:"总数"`
}
