package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AddProductCategoryReq struct {
	g.Meta `path:"/product-category" tags:"产品分类管理" method:"post" summary:"新增"`
	model.Author
	PID   int64  `json:"pid" v:"required#产品分类父ID不能为空" dc:"父ID"`
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name  string `json:"name" v:"required#产品分类名称不能为空" dc:"名称"`
}

type AddProductCategoryRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"分类ID"`
}

type DeleteProductCategoryReq struct {
	g.Meta `path:"/product-category/{id}" tags:"产品分类管理" method:"delete" summary:"删除"`
	model.Author
	ID int64 `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type DeleteProductCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateProductCategoryReq struct {
	g.Meta `path:"/product-category/{id}" tags:"产品分类管理" method:"put" summary:"修改(全量更新)"`
	model.Author
	ID    int64  `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
	PID   int64  `json:"pid" v:"required#父级分类ID不能为空" dc:"父级分类ID"`
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name  string `json:"name" v:"required#产品分类名称不能为空" dc:"名称"`
}

type UpdateProductCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type GetProductCategoryTreeByIDReq struct {
	g.Meta `path:"/product-category/{id}" tags:"产品分类管理" method:"get" summary:"根据ID查询"`
	model.Author
	ID int64 `p:"id" v:"required#产品分类ID不能为空" dc:"产品分类ID"`
}

type GetProductCategoryTreeByIDRes struct {
	g.Meta `mime:"application/json"`
	*model.ProductCategoryTree
}

type GetProductCategoryTreeByOrgIDReq struct {
	g.Meta `path:"/product-category/tree" tags:"产品分类管理" method:"get" summary:"根据组织ID查询"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type GetProductCategoryTreeByOrgIDRes struct {
	g.Meta                  `mime:"application/json"`
	ProductCategoryTreeList model.ProductCategoryTreeList `json:"list" dc:"分类列表"`
}
