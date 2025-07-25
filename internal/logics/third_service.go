package logics

import (
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"sync"
)

var (
	thirdServiceOnce     sync.Once
	thirdServiceInstance *thirdService
)

type thirdService struct {
	IntrospectAddr string
}

func NewThirdService() service.IThirdService {
	thirdServiceOnce.Do(func() {
		thirdServiceInstance = &thirdService{
			IntrospectAddr: "http://127.0.0.1:8080/api/v1/identify-service/token/introspect",
		}
	})
	return thirdServiceInstance
}

func (s *thirdService) Introspect(ctx context.Context, token string) (out *model.IntrospectRes, err error) {
	return
}
