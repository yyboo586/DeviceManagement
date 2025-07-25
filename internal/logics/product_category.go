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

	"github.com/gogf/gf/v2/frame/g"
)

var (
	productCategoryOnce     sync.Once
	productCategoryInstance *productCategory
)

type productCategory struct {
}

func NewProductCategory() service.IProductCategory {
	productCategoryOnce.Do(func() {
		productCategoryInstance = &productCategory{}
	})
	return productCategoryInstance
}

func (l *productCategory) Add(ctx context.Context, req *v1.AddProductCategoryReq) (id int64, err error) {
	dataInsert := map[string]interface{}{
		dao.ProductCategory.Columns().OrgID: req.OrgID,
		dao.ProductCategory.Columns().Name:  req.Name,
		dao.ProductCategory.Columns().Desc:  req.Desc,
	}

	id, err = dao.ProductCategory.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return id, nil
}

func (l *productCategory) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.ProductCategory.Ctx(ctx).WhereIn(dao.ProductCategory.Columns().ID, ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return
}

func (l *productCategory) Update(ctx context.Context, req *v1.UpdateProductCategoryReq) (err error) {
	dataUpdate := map[string]interface{}{
		dao.ProductCategory.Columns().Name: req.Name,
		dao.ProductCategory.Columns().Desc: req.Desc,
	}

	_, err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().ID, req.ID).Data(dataUpdate).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return
}

func (l *productCategory) Get(ctx context.Context, id int64) (out *model.ProductCategory, err error) {
	var entity entity.TProductCategory
	err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("分类不存在")
		}
		return nil, err
	}

	out = l.convertEntityToModel(&entity)
	return
}

func (l *productCategory) List(ctx context.Context, in *v1.ListProductCategoryReq) (out []*model.ProductCategory, total int, currentPage int, err error) {
	if in.PageReq.Page <= 0 {
		in.PageReq.Page = 1
		currentPage = 1
	}
	if in.PageReq.PageSize <= 0 {
		in.PageReq.PageSize = model.DefaultPageSize
	}

	m := dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().OrgID, in.OrgID)
	if in.Name != "" {
		m = m.Where(dao.ProductCategory.Columns().Name, in.Name)
	}

	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var result []*entity.TProductCategory
	err = m.Page(in.PageReq.Page, in.PageReq.PageSize).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range result {
		out = append(out, l.convertEntityToModel(v))
	}
	return
}

func (l *productCategory) convertEntityToModel(in *entity.TProductCategory) (out *model.ProductCategory) {
	out = &model.ProductCategory{
		ID:        in.ID,
		OrgID:     in.OrgID,
		Name:      in.Name,
		Desc:      in.Desc,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
