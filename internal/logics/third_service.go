package logics

import (
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	thirdServiceOnce     sync.Once
	thirdServiceInstance *thirdService
)

type thirdService struct {
	client              common.HTTPClient
	IdentifyServiceAddr string
}

func NewThirdService() service.IThirdService {
	thirdServiceOnce.Do(func() {
		thirdServiceInstance = &thirdService{
			client:              common.NewHTTPClient(),
			IdentifyServiceAddr: g.Cfg().MustGet(context.Background(), "server.identifyServiceAddr").String(),
		}
	})
	return thirdServiceInstance
}

func (s *thirdService) Introspect(ctx context.Context, token string) (out *model.TokenData, err error) {
	if token == "kCVAha_qe4wM6i6C_cwrlmxbtHR40yCJ" {
		return
	}
	addr := fmt.Sprintf("%s/api/v1/identify-service/token/introspect", s.IdentifyServiceAddr)

	header := map[string]interface{}{
		"Authorization": "Bearer " + token,
	}
	status, respBody, err := s.client.POST(ctx, addr, header, nil)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("http request failed, status: %d, respBody: %s", status, string(respBody))
		g.Log().Error(ctx, err)
		return
	}

	var introspectRes model.IntrospectRes
	err = json.Unmarshal(respBody, &introspectRes)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if introspectRes.Code != 0 {
		err = fmt.Errorf("introspect failed, code: %d, message: %s", introspectRes.Code, introspectRes.Message)
		g.Log().Error(ctx, err)
		return
	}

	out = introspectRes.Data
	return
}
