package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IProduct interface {
	Add(ctx context.Context, req *v1.AddProductReq) (out *model.Product, err error)

	Delete(ctx context.Context, id int64) (err error)

	Update(ctx context.Context, req *v1.UpdateProductReq) (err error)

	Get(ctx context.Context, id int64) (out *model.Product, err error)
	// 获取某组织下所有产品列表，或者某分类下所有产品列表
	List(ctx context.Context, req *v1.ListProductReq) (out []*model.Product, err error)
}

var (
	localProduct IProduct
)

func Product() IProduct {
	if localProduct == nil {
		panic("implement not found for interface IProduct, forgot register?")
	}
	return localProduct
}

func RegisterProduct(i IProduct) {
	localProduct = i
}
