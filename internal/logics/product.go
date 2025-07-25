package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"DeviceManagement/internal/service"
	"context"
	"database/sql"
	"errors"
	"sync"
)

var (
	productOnce     sync.Once
	productInstance *product
)

type product struct {
}

func NewProduct() service.IProduct {
	productOnce.Do(func() {
		productInstance = &product{}
	})
	return productInstance
}

func (l *product) Add(ctx context.Context, req *v1.AddProductReq) (out *model.Product, err error) {
	dataInsert := map[string]interface{}{
		dao.Product.Columns().CategoryID: req.CategoryID,
		dao.Product.Columns().OrgID:      req.OrgID,
		dao.Product.Columns().Name:       req.Name,
	}

	id, err := dao.Product.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	out = &model.Product{
		ID: id,
	}
	return
}

func (l *product) Delete(ctx context.Context, id int64) (err error) {
	_, err = dao.Product.Ctx(ctx).Where(dao.Product.Columns().ID, id).Delete()

	return
}

func (l *product) Update(ctx context.Context, req *v1.UpdateProductReq) (err error) {
	dataUpdate := map[string]interface{}{
		dao.Product.Columns().CategoryID: req.CategoryID,
		dao.Product.Columns().Name:       req.Name,
	}

	_, err = dao.Product.Ctx(ctx).Where(dao.Product.Columns().ID, req.ID).Data(dataUpdate).Update()

	return
}

func (l *product) Get(ctx context.Context, id int64) (out *model.Product, err error) {
	var entity entity.TProduct
	err = dao.Product.Ctx(ctx).Where(dao.Product.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("产品不存在")
		}
		return nil, err
	}

	out = l.convertEntityToLogics(&entity)
	return
}

func (l *product) List(ctx context.Context, req *v1.ListProductReq) (out []*model.Product, err error) {
	model := dao.Product.Ctx(ctx).Where(dao.Product.Columns().OrgID, req.OrgID)
	// 分类ID过滤
	if req.CategoryID > 0 {
		model = model.Where(dao.Product.Columns().CategoryID, req.CategoryID)
	}

	var result []*entity.TProduct
	err = model.Scan(&result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return
	}

	for _, entity := range result {
		out = append(out, l.convertEntityToLogics(entity))
	}
	return
}

func (l *product) convertEntityToLogics(in *entity.TProduct) (out *model.Product) {
	return &model.Product{
		ID:         in.ID,
		OrgID:      in.OrgID,
		Name:       in.Name,
		CategoryID: in.CategoryID,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}
}
