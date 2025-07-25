package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/service"
	"context"
)

var (
	ProductController = &productController{}
)

type productController struct {
}

func (c *productController) Add(ctx context.Context, req *v1.AddProductReq) (res *v1.AddProductRes, err error) {
	out, err := service.Product().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.AddProductRes{
		ID: out.ID,
	}
	return
}

func (c *productController) Update(ctx context.Context, req *v1.UpdateProductReq) (res *v1.UpdateProductRes, err error) {
	err = service.Product().Update(ctx, req)
	return
}

func (c *productController) Delete(ctx context.Context, req *v1.DeleteProductReq) (res *v1.DeleteProductRes, err error) {
	err = service.Product().Delete(ctx, req.ID)
	return
}

func (c *productController) Get(ctx context.Context, req *v1.GetProductReq) (res *v1.GetProductRes, err error) {
	out, err := service.Product().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetProductRes{
		Product: out,
	}
	return
}

func (c *productController) List(ctx context.Context, req *v1.ListProductReq) (res *v1.ListProductRes, err error) {
	out, err := service.Product().List(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.ListProductRes{}
	res.List = append(res.List, out...)
	res.Total = len(out)
	return
}
