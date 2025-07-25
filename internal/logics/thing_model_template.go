package logics

import (
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
	thingModelTemplateOnce     sync.Once
	thingModelTemplateInstance *thingModelTemplate
)

type thingModelTemplate struct{}

func NewThingModelTemplate() service.IThingModelTemplate {
	thingModelTemplateOnce.Do(func() {
		thingModelTemplateInstance = &thingModelTemplate{}
	})
	return thingModelTemplateInstance
}

func (l *thingModelTemplate) Get(ctx context.Context, id int64) (out *model.ThingModelTemplate, err error) {
	var entity *entity.TThingModelTemplate
	err = dao.ThingModelTemplate.Ctx(ctx).Where(dao.ThingModelTemplate.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("模板不存在")
		}
		return nil, err
	}

	out, err = l.convertEntityToLogics(entity)

	return
}

func (l *thingModelTemplate) List(ctx context.Context, orgID string) (out []*model.ThingModelTemplate, err error) {
	var entities []*entity.TThingModelTemplate
	// 获取系统内置模板
	err = dao.ThingModelTemplate.Ctx(ctx).Where(dao.ThingModelTemplate.Columns().OrgID, "").Scan(&entities)
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		template, err := l.convertEntityToLogics(entity)
		if err != nil {
			return nil, err
		}
		out = append(out, template)
	}

	// 获取组织模板
	entities = make([]*entity.TThingModelTemplate, 0)
	err = dao.ThingModelTemplate.Ctx(ctx).Where(dao.ThingModelTemplate.Columns().OrgID, orgID).Scan(&entities)
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		template, err := l.convertEntityToLogics(entity)
		if err != nil {
			return nil, err
		}
		out = append(out, template)
	}
	return
}

func (l *thingModelTemplate) convertEntityToLogics(in *entity.TThingModelTemplate) (out *model.ThingModelTemplate, err error) {
	properties := make([]*model.ThingModelProperty, 0)
	if err = json.Unmarshal([]byte(in.Properties), &properties); err != nil {
		return nil, err
	}
	services := make([]*model.ThingModelService, 0)
	if err = json.Unmarshal([]byte(in.Services), &services); err != nil {
		return nil, err
	}
	events := make([]*model.ThingModelEvent, 0)
	if err = json.Unmarshal([]byte(in.Events), &events); err != nil {
		return nil, err
	}

	out = &model.ThingModelTemplate{
		ID:          in.ID,
		OrgID:       in.OrgID,
		Name:        in.Name,
		Description: in.Description,
		Properties:  properties,
		Services:    services,
		Events:      events,
		IsSystem:    in.IsSystem == 0,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return
}
