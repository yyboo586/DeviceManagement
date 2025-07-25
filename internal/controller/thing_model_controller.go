package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/service"
	"context"
)

var (
	ThingModelController = &thingModelController{}
)

type thingModelController struct {
}

func (c *thingModelController) CreateFromTemplate(ctx context.Context, req *v1.CreateFromTemplateReq) (res *v1.CreateFromTemplateRes, err error) {
	id, err := service.ThingModel().CreateFromTemplate(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateFromTemplateRes{
		ID: id,
	}
	return
}

func (c *thingModelController) Update(ctx context.Context, req *v1.UpdateThingModelReq) (res *v1.UpdateThingModelRes, err error) {
	err = service.ThingModel().Update(ctx, req)
	return
}

func (c *thingModelController) Delete(ctx context.Context, req *v1.DeleteThingModelReq) (res *v1.DeleteThingModelRes, err error) {
	err = service.ThingModel().Delete(ctx, req.ID)
	return
}

func (c *thingModelController) Get(ctx context.Context, req *v1.GetThingModelReq) (res *v1.GetThingModelRes, err error) {
	out, err := service.ThingModel().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetThingModelRes{
		ThingModel: out,
	}
	return
}

func (c *thingModelController) List(ctx context.Context, req *v1.ListThingModelReq) (res *v1.ListThingModelRes, err error) {
	out, err := service.ThingModel().List(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	res = &v1.ListThingModelRes{}
	res.List = append(res.List, out...)
	res.Total = len(out)
	return
}

func (c *thingModelController) GetBaseTemplates(ctx context.Context, req *v1.GetThingModelTemplatesReq) (res *v1.GetThingModelTemplatesRes, err error) {
	templates, err := service.ThingModelTemplate().List(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	res = &v1.GetThingModelTemplatesRes{
		Templates: templates,
	}
	return
}
