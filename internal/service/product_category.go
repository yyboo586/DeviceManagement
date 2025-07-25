package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IProductCategory interface {
	Add(ctx context.Context, req *v1.AddProductCategoryReq) (out *model.ProductCategory, err error)

	Delete(ctx context.Context, req *v1.DeleteProductCategoryReq) (err error)

	Update(ctx context.Context, req *v1.UpdateProductCategoryReq) (err error)

	GetTreeByID(ctx context.Context, id int64) (out *model.ProductCategoryTree, err error)
	GetTreeByOrgID(ctx context.Context, orgID string) (out model.ProductCategoryTreeList, err error)
}

var (
	localProductCategory IProductCategory
)

func ProductCategory() IProductCategory {
	if localProductCategory == nil {
		panic("implement not found for interface IProductCategory, forgot register?")
	}
	return localProductCategory
}

func RegisterProductCategory(i IProductCategory) {
	localProductCategory = i
}
