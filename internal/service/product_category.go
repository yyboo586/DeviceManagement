package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IProductCategory interface {
	Add(ctx context.Context, in *v1.AddProductCategoryReq) (id int64, err error)

	Delete(ctx context.Context, ids []int64) (err error)

	Update(ctx context.Context, in *v1.UpdateProductCategoryReq) (err error)

	// 详情
	Get(ctx context.Context, id int64) (out *model.ProductCategory, err error)
	// 列表
	List(ctx context.Context) (out []*model.ProductCategory, err error)
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
