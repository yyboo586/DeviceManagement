package service

import (
	"DeviceManagement/internal/model"
	"context"
)

type IThingModelTemplate interface {
	// 获取物模型模板详情
	Get(ctx context.Context, id int64) (out *model.ThingModelTemplate, err error)
	// 获取物模型模板列表
	List(ctx context.Context, orgID string) (out []*model.ThingModelTemplate, err error)
}

var (
	localThingModelTemplate IThingModelTemplate
)

func ThingModelTemplate() IThingModelTemplate {
	if localThingModelTemplate == nil {
		panic("implement not found for interface IThingModelTemplate, forgot register?")
	}
	return localThingModelTemplate
}

func RegisterThingModelTemplate(i IThingModelTemplate) {
	localThingModelTemplate = i
}
