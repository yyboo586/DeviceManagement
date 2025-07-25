package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddProductCategoryReq struct {
	g.Meta `path:"/product-category" tags:"分类管理" method:"post" summary:"添加"`
	model.Author
	Name string `json:"name" v:"required#产品分类名称不能为空" dc:"产品分类名称"`
	Desc string `json:"desc" dc:"产品分类描述"`
}

type AddProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"产品分类ID"`
}

type DeleteProductCategoryReq struct {
	g.Meta `path:"/product-category" tags:"分类管理" method:"delete" summary:"删除"`
	model.Author
	IDs []int64 `json:"ids" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type DeleteProductCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateProductCategoryReq struct {
	g.Meta `path:"/product-category/{id}" tags:"分类管理" method:"put" summary:"编辑"`
	model.Author
	ID   int64  `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	Name string `json:"name" v:"required#产品分类名称不能为空" dc:"产品分类名称"`
	Desc string `json:"desc" dc:"产品分类描述"`
}

type UpdateProductCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type GetProductCategoryReq struct {
	g.Meta `path:"/product-category/{id}" tags:"分类管理" method:"get" summary:"详情"`
	model.Author
	ID int64 `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type GetProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	*ProductCategory
}

type ListProductCategoryReq struct {
	g.Meta `path:"/product-category" tags:"分类管理" method:"get" summary:"列表"`
	model.Author
}

type ListProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	List   []*ProductCategory `json:"list" dc:"列表"`
}

type GetProductCategoryTreeReq struct {
	g.Meta `path:"/product-category/{id}/tree" tags:"分类管理" method:"get" summary:"详情(树形结构)"`
	model.Author
	ID int64 `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type GetProductCategoryTreeRes struct {
	g.Meta `mime:"application/json"`
	*ProductCategory
}

type ListProductCategoryTreeReq struct {
	g.Meta `path:"/product-category/trees" tags:"分类管理" method:"get" summary:"列表(树形结构)"`
}

type ListProductCategoryTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*ProductCategory `json:"list" dc:"列表"`
}

type ProductCategory struct {
	ID        int64       `json:"id" dc:"产品分类ID"`
	Name      string      `json:"name" dc:"产品分类名称"`
	Desc      string      `json:"desc" dc:"产品分类描述"`
	Products  []*Product  `json:"products" dc:"产品列表"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}
