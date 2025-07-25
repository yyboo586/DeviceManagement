package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IProduct interface {
	Add(ctx context.Context, req *v1.AddProductReq) (id int64, err error)

	Delete(ctx context.Context, ids []int64) (err error)

	Update(ctx context.Context, req *v1.UpdateProductReq) (err error)

	Get(ctx context.Context, id int64) (out *model.Product, err error)
	List(ctx context.Context, req *v1.ListProductReq) (out []*model.Product, pageRes *model.PageRes, err error)
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
