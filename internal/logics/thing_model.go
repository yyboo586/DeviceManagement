package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"DeviceManagement/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"sync"
)

var (
	thingModelOnce     sync.Once
	thingModelInstance *thingModel
)

type thingModel struct {
}

func NewThingModel() service.IThingModel {
	thingModelOnce.Do(func() {
		thingModelInstance = &thingModel{}
	})
	return thingModelInstance
}

// 从模板创建物模型
func (l *thingModel) CreateFromTemplate(ctx context.Context, req *v1.CreateFromTemplateReq) (id int64, err error) {
	// 获取模板
	template, err := thingModelTemplateInstance.Get(ctx, req.TemplateID)
	if err != nil {
		return
	}

	// 创建物模型
	dataInsert := map[string]interface{}{
		dao.ThingModel.Columns().OrgID:       req.OrgID,
		dao.ThingModel.Columns().ProductID:   req.ProductID,
		dao.ThingModel.Columns().TemplateID:  req.TemplateID,
		dao.ThingModel.Columns().Name:        req.Name,
		dao.ThingModel.Columns().Description: req.Description,
		dao.ThingModel.Columns().Properties:  template.Properties,
		dao.ThingModel.Columns().Services:    template.Services,
		dao.ThingModel.Columns().Events:      template.Events,
	}

	id, err = dao.ThingModel.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		return
	}

	return
}

func (l *thingModel) Delete(ctx context.Context, id int64) (err error) {
	_, err = dao.ThingModel.Ctx(ctx).Where(dao.ThingModel.Columns().ID, id).Delete()

	return
}

func (l *thingModel) Update(ctx context.Context, req *v1.UpdateThingModelReq) (err error) {
	template, err := thingModelTemplateInstance.Get(ctx, req.TemplateID)
	if err != nil {
		return
	}

	dataUpdate := map[string]interface{}{
		dao.ThingModel.Columns().OrgID:       req.OrgID,
		dao.ThingModel.Columns().ProductID:   req.ProductID,
		dao.ThingModel.Columns().TemplateID:  req.TemplateID,
		dao.ThingModel.Columns().Name:        req.Name,
		dao.ThingModel.Columns().Version:     req.Version,
		dao.ThingModel.Columns().Description: req.Description,
		dao.ThingModel.Columns().Properties:  template.Properties,
		dao.ThingModel.Columns().Services:    template.Services,
		dao.ThingModel.Columns().Events:      template.Events,
	}

	_, err = dao.ThingModel.Ctx(ctx).Where(dao.ThingModel.Columns().ID, req.ID).Data(dataUpdate).Update()

	return
}

func (l *thingModel) Get(ctx context.Context, id int64) (out *model.ThingModel, err error) {
	var entity entity.TThingModel
	err = dao.ThingModel.Ctx(ctx).Where(dao.ThingModel.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("物模型不存在")
		}
		return nil, err
	}

	out, err = l.convertEntityToLogics(&entity)
	if err != nil {
		return nil, err
	}

	return
}

func (l *thingModel) List(ctx context.Context, orgID string) (out []*model.ThingModel, err error) {
	var result []*entity.TThingModel
	err = dao.ThingModel.Ctx(ctx).Where(dao.ThingModel.Columns().OrgID, orgID).Scan(&result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return
	}

	for _, entity := range result {
		thingModel, err := l.convertEntityToLogics(entity)
		if err != nil {
			return nil, err
		}
		out = append(out, thingModel)
	}

	return
}

func (l *thingModel) convertEntityToLogics(in *entity.TThingModel) (out *model.ThingModel, err error) {
	properties := make([]*model.ThingModelProperty, 0)
	err = json.Unmarshal([]byte(in.Properties), &properties)
	if err != nil {
		return nil, err
	}
	services := make([]*model.ThingModelService, 0)
	err = json.Unmarshal([]byte(in.Services), &services)
	if err != nil {
		return nil, err
	}
	events := make([]*model.ThingModelEvent, 0)
	err = json.Unmarshal([]byte(in.Events), &events)
	if err != nil {
		return nil, err
	}

	out = &model.ThingModel{
		ID:          in.ID,
		ProductID:   in.ProductID,
		OrgID:       in.OrgID,
		TemplateID:  in.TemplateID,
		Name:        in.Name,
		Version:     in.Version,
		Description: in.Description,
		Properties:  properties,
		Services:    services,
		Events:      events,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return out, nil
}
