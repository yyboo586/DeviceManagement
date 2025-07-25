package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/service"
	"context"
)

var ProductCategoryController = &productCategoryController{}

type productCategoryController struct {
}

func (c *productCategoryController) Add(ctx context.Context, req *v1.AddProductCategoryReq) (res *v1.AddProductCategoryRes, err error) {
	out, err := service.ProductCategory().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.AddProductCategoryRes{
		ID: out.ID,
	}
	return
}

func (c *productCategoryController) Delete(ctx context.Context, req *v1.DeleteProductCategoryReq) (res *v1.DeleteProductCategoryRes, err error) {
	err = service.ProductCategory().Delete(ctx, req)
	return
}

func (c *productCategoryController) Update(ctx context.Context, req *v1.UpdateProductCategoryReq) (res *v1.UpdateProductCategoryRes, err error) {
	err = service.ProductCategory().Update(ctx, req)
	return
}

func (c *productCategoryController) GetTreeByID(ctx context.Context, req *v1.GetProductCategoryTreeByIDReq) (res *v1.GetProductCategoryTreeByIDRes, err error) {
	out, err := service.ProductCategory().GetTreeByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetProductCategoryTreeByIDRes{
		ProductCategoryTree: out,
	}
	return
}

func (c *productCategoryController) GetTreeByOrgID(ctx context.Context, req *v1.GetProductCategoryTreeByOrgIDReq) (res *v1.GetProductCategoryTreeByOrgIDRes, err error) {
	out, err := service.ProductCategory().GetTreeByOrgID(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetProductCategoryTreeByOrgIDRes{
		ProductCategoryTreeList: out,
	}
	return
}
