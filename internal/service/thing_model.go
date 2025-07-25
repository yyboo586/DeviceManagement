package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type IThingModel interface {
	// 从模板创建物模型
	CreateFromTemplate(ctx context.Context, req *v1.CreateFromTemplateReq) (id int64, err error)

	// 删除物模型
	Delete(ctx context.Context, id int64) (err error)

	// 更新物模型
	Update(ctx context.Context, req *v1.UpdateThingModelReq) (err error)

	// 获取物模型详情
	Get(ctx context.Context, id int64) (out *model.ThingModel, err error)
	// 获取物模型列表
	List(ctx context.Context, orgID string) (out []*model.ThingModel, err error)
}

var (
	localThingModel IThingModel
)

func ThingModel() IThingModel {
	if localThingModel == nil {
		panic("implement not found for interface IThingModel, forgot register?")
	}
	return localThingModel
}

func RegisterThingModel(i IThingModel) {
	localThingModel = i
}
