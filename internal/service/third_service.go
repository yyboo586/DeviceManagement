package service

import (
	"DeviceManagement/internal/model"
	"context"
)

type IThirdService interface {
	Introspect(ctx context.Context, token string) (out *model.IntrospectRes, err error)
}

var (
	localThirdService IThirdService
)

func ThirdService() IThirdService {
	if localThirdService == nil {
		panic("implement not found for interface IThirdService, forgot register?")
	}
	return localThirdService
}

func RegisterThirdService(i IThirdService) {
	localThirdService = i
}
