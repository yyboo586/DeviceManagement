package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/service"
	"context"
)

var CronJobTemplateController = cronJobTemplateController{}

type cronJobTemplateController struct {
}

func (c *cronJobTemplateController) List(ctx context.Context, req *v1.ListCronJobTemplateReq) (res *v1.ListCronJobTemplateRes, err error) {
	out, pageRes, err := service.CronJobTemplate().List(ctx, req)
	if err != nil {
		return
	}

	res = &v1.ListCronJobTemplateRes{
		List:    out,
		PageRes: *pageRes,
	}
	return
}
