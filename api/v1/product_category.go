package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddProductCategoryReq struct {
	g.Meta `path:"/product-category" tags:"产品分类" method:"post" summary:"新增"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name  string `json:"name" v:"required#产品分类名称不能为空" dc:"产品分类名称"`
	Desc  string `json:"desc" dc:"产品分类描述"`
}

type AddProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"产品分类ID"`
}

type DeleteProductCategoryReq struct {
	g.Meta `path:"/product-category" tags:"产品分类" method:"delete" summary:"删除"`
	model.Author
	IDs []int64 `json:"ids" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type DeleteProductCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateProductCategoryReq struct {
	g.Meta `path:"/product-category/{id}" tags:"产品分类" method:"put" summary:"修改(全量更新)"`
	model.Author
	ID   int64  `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	Name string `json:"name" v:"required#产品分类名称不能为空" dc:"产品分类名称"`
	Desc string `json:"desc" dc:"产品分类描述"`
}

type UpdateProductCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type GetProductCategoryReq struct {
	g.Meta `path:"/product-category/{id}" tags:"产品分类" method:"get" summary:"详情"`
	model.Author
	ID int64 `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type GetProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	*ProductCategory
}

type ListProductCategoryReq struct {
	g.Meta `path:"/product-category/trees" tags:"产品分类" method:"get" summary:"树形列表"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name  string `json:"name" dc:"产品分类名称"`
	model.PageReq
}

type ListProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	List   []*ProductCategory `json:"list" dc:"分类列表"`
	model.PageRes
}

type ProductCategory struct {
	ID        int64       `json:"id" dc:"产品分类ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	Name      string      `json:"name" dc:"产品分类名称"`
	Desc      string      `json:"desc" dc:"产品分类描述"`
	Products  []*Product  `json:"products" dc:"产品列表"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}
