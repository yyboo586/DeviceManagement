package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

var (
	ProductController = &productController{}
)

type productController struct {
}

func (c *productController) Add(ctx context.Context, req *v1.AddProductReq) (res *v1.AddProductRes, err error) {
	id, err := service.Product().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.AddProductRes{
		ID: id,
	}
	return
}

func (c *productController) Update(ctx context.Context, req *v1.UpdateProductReq) (res *v1.UpdateProductRes, err error) {
	err = service.Product().Update(ctx, req)
	return
}

func (c *productController) Delete(ctx context.Context, req *v1.DeleteProductReq) (res *v1.DeleteProductRes, err error) {
	err = service.Product().Delete(ctx, req.IDs)
	return
}

func (c *productController) Get(ctx context.Context, req *v1.GetProductReq) (res *v1.GetProductRes, err error) {
	out, err := service.Product().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetProductRes{
		Product: c.convert(out),
	}
	return
}

func (c *productController) List(ctx context.Context, req *v1.ListProductReq) (res *v1.ListProductRes, err error) {
	out, pageRes, err := service.Product().List(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.ListProductRes{
		PageRes: pageRes,
	}
	for _, v := range out {
		res.List = append(res.List, c.convert(v))
	}
	return
}

func (c *productController) convert(in *model.Product) (out *v1.Product) {
	out = &v1.Product{
		ID:         in.ID,
		CategoryID: in.CategoryID,
		Name:       in.Name,
		Desc:       in.Desc,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}
	return
}
