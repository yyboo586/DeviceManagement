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
	req.OrgID = "00000000-0000-0000-0000-000000000000"
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
	req.OrgID = "00000000-0000-0000-0000-000000000000"
	out, total, currentPage, err := service.ProductCategory().List(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.ListProductCategoryRes{
		PageRes: model.PageRes{
			Total:       total,
			CurrentPage: currentPage,
		},
	}
	for _, v := range out {
		out, _, _, err := service.Product().List(ctx, &v1.ListProductReq{
			OrgID:      req.OrgID,
			CategoryID: v.ID,
		})
		if err != nil {
			return nil, err
		}
		item := c.convert(v)
		for _, product := range out {
			item.Products = append(item.Products, ProductController.convert(product))
		}
		res.List = append(res.List, item)
	}

	return
}

func (c *productCategoryController) convert(in *model.ProductCategory) *v1.ProductCategory {
	return &v1.ProductCategory{
		ID:        in.ID,
		OrgID:     in.OrgID,
		Name:      in.Name,
		Desc:      in.Desc,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
