package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

var ProductCategoryController = &productCategoryController{}

type productCategoryController struct {
}

func (c *productCategoryController) Add(ctx context.Context, req *v1.AddProductCategoryReq) (res *v1.AddProductCategoryRes, err error) {
	id, err := service.ProductCategory().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.AddProductCategoryRes{
		ID: id,
	}
	return
}

func (c *productCategoryController) Delete(ctx context.Context, req *v1.DeleteProductCategoryReq) (res *v1.DeleteProductCategoryRes, err error) {
	err = service.ProductCategory().Delete(ctx, req.IDs)
	return
}

func (c *productCategoryController) Update(ctx context.Context, req *v1.UpdateProductCategoryReq) (res *v1.UpdateProductCategoryRes, err error) {
	err = service.ProductCategory().Update(ctx, req)
	return
}

func (c *productCategoryController) Get(ctx context.Context, req *v1.GetProductCategoryReq) (res *v1.GetProductCategoryRes, err error) {
	out, err := service.ProductCategory().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetProductCategoryRes{
		ProductCategory: c.convert(out),
	}
	return
}

func (c *productCategoryController) List(ctx context.Context, req *v1.ListProductCategoryReq) (res *v1.ListProductCategoryRes, err error) {
	out, err := service.ProductCategory().List(ctx)
	if err != nil {
		return
	}

	res = &v1.ListProductCategoryRes{}
	for _, v := range out {
		item := c.convert(v)
		res.List = append(res.List, item)
	}

	return
}

func (c *productCategoryController) GetTree(ctx context.Context, req *v1.GetProductCategoryTreeReq) (res *v1.GetProductCategoryTreeRes, err error) {
	out, err := service.ProductCategory().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetProductCategoryTreeRes{
		ProductCategory: c.convert(out),
	}

	products, _, err := service.Product().List(ctx, &v1.ListProductReq{
		CategoryID: req.ID,
	})
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		res.ProductCategory.Products = append(res.ProductCategory.Products, ProductController.convert(product))
	}
	return
}

func (c *productCategoryController) ListTree(ctx context.Context, req *v1.ListProductCategoryTreeReq) (res *v1.ListProductCategoryTreeRes, err error) {
	out, err := service.ProductCategory().List(ctx)
	if err != nil {
		return
	}

	res = &v1.ListProductCategoryTreeRes{}
	for _, v := range out {
		products, _, err := service.Product().List(ctx, &v1.ListProductReq{
			CategoryID: v.ID,
		})
		if err != nil {
			return nil, err
		}
		item := c.convert(v)
		for _, product := range products {
			item.Products = append(item.Products, ProductController.convert(product))
		}
		res.List = append(res.List, item)
	}

	return
}

func (c *productCategoryController) convert(in *model.ProductCategory) *v1.ProductCategory {
	return &v1.ProductCategory{
		ID:        in.ID,
		Name:      in.Name,
		Desc:      in.Desc,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
