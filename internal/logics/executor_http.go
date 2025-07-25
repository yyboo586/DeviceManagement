package logics

import (
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/model"
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/os/gtime"
)

// HTTPExecutor HTTP任务执行器
type HTTPExecutor struct {
	client common.HTTPClient
}

var (
	httpExecutorOnce     sync.Once
	httpExecutorInstance *HTTPExecutor
)

func NewHTTPExecutor() *HTTPExecutor {
	httpExecutorOnce.Do(func() {
		httpExecutorInstance = &HTTPExecutor{
			client: common.NewHTTPClient(),
		}
	})
	return httpExecutorInstance
}

func (e *HTTPExecutor) Execute(ctx context.Context, params interface{}) *model.ExecutionResult {
	result := &model.ExecutionResult{
		StartTime: gtime.Now(),
	}

	config := &model.HTTPRequestConfig{
		Method: "GET",
		URL:    fmt.Sprintf("http://127.0.0.1:9501/api/v1/device-management/devices?org_id=%s", "1"),
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
	}

	// 执行HTTP请求
	status, respBody, err := e.executeHTTPRequest(ctx, config)

	result.EndTime = gtime.Now()
	result.Duration = result.EndTime.Sub(result.StartTime).Milliseconds()

	if err != nil {
		result.Success = false
		result.Message = "HTTP请求失败: " + err.Error()
		result.Error = err
		return result
	}

	// 判断响应状态
	if status >= 200 && status < 300 {
		result.Success = true
		result.Message = "HTTP请求执行成功"
		result.Data = map[string]interface{}{
			"status": status,
			"body":   string(respBody),
		}
	} else {
		result.Success = false
		result.Message = fmt.Sprintf("HTTP请求失败, 状态码: %d", status)
		result.Data = map[string]interface{}{
			"status": status,
			"body":   string(respBody),
		}
	}

	return result
}

// executeHTTPRequest 执行HTTP请求
func (e *HTTPExecutor) executeHTTPRequest(ctx context.Context, config *model.HTTPRequestConfig) (status int, respBody []byte, err error) {
	switch config.Method {
	case "GET":
		return e.client.GET(ctx, config.URL, config.Headers)
	case "POST":
		return e.client.POST(ctx, config.URL, config.Headers, config.Body)
	default:
		err = fmt.Errorf("不支持的HTTP方法: %s", config.Method)
		return
	}
}
