package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"DeviceManagement/internal/service"
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
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

func (l *product) Add(ctx context.Context, req *v1.AddProductReq) (id int64, err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	dataInsert := map[string]interface{}{
		dao.Product.Columns().OrgID:      operatorInfo.OrgID,
		dao.Product.Columns().CategoryID: req.CategoryID,
		dao.Product.Columns().Name:       req.Name,
		dao.Product.Columns().Desc:       req.Desc,
	}

	id, err = dao.Product.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (l *product) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.Product.Ctx(ctx).WhereIn(dao.Product.Columns().ID, ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return
}

func (l *product) Update(ctx context.Context, req *v1.UpdateProductReq) (err error) {
	dataUpdate := map[string]interface{}{
		dao.Product.Columns().CategoryID: req.CategoryID,
		dao.Product.Columns().Name:       req.Name,
		dao.Product.Columns().Desc:       req.Desc,
	}

	_, err = dao.Product.Ctx(ctx).Where(dao.Product.Columns().ID, req.ID).Data(dataUpdate).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
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

	out = l.convertEntityToModel(&entity)
	return
}

func (l *product) List(ctx context.Context, req *v1.ListProductReq) (out []*model.Product, pageRes *model.PageRes, err error) {
	if req.PageReq.Page <= 0 {
		req.PageReq.Page = 1
	}
	if req.PageReq.PageSize <= 0 {
		req.PageReq.PageSize = model.DefaultPageSize
	}

	m := dao.Product.Ctx(ctx)
	// 分类ID过滤
	if req.CategoryID > 0 {
		m = m.Where(dao.Product.Columns().CategoryID, req.CategoryID)
	}

	total, err := m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var result []*entity.TProduct
	err = m.Page(req.PageReq.Page, req.PageReq.PageSize).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	for _, entity := range result {
		out = append(out, l.convertEntityToModel(entity))
	}
	pageRes = &model.PageRes{
		CurrentPage: req.PageReq.Page,
		Total:       total,
	}

	return
}

func (l *product) convertEntityToModel(in *entity.TProduct) (out *model.Product) {
	return &model.Product{
		ID:         in.ID,
		OrgID:      in.OrgID,
		CategoryID: in.CategoryID,
		Name:       in.Name,
		Desc:       in.Desc,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}
}
