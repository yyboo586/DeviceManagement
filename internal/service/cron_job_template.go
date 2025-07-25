package service

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"context"
)

type ICronJobTemplate interface {
	List(ctx context.Context, in *v1.ListCronJobTemplateReq) (out []*model.CronJobTemplate, pageRes *model.PageRes, err error)
}

var localCronJobTemplate ICronJobTemplate

func CronJobTemplate() ICronJobTemplate {
	if localCronJobTemplate == nil {
		panic("implement not found for interface ICronJobTemplate, forgot register?")
	}
	return localCronJobTemplate
}

func RegisterCronJobTemplate(i ICronJobTemplate) {
	localCronJobTemplate = i
}
