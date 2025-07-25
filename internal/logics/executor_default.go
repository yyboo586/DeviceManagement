package logics

import (
	"DeviceManagement/internal/model"
	"context"
	"sync"
)

// DefaultExecutor 默认任务执行器
type DefaultExecutor struct {
	httpExecutor *HTTPExecutor
}

var (
	defaultExecutorOnce     sync.Once
	defaultExecutorInstance *DefaultExecutor
)

func NewDefaultExecutor() *DefaultExecutor {
	defaultExecutorOnce.Do(func() {
		defaultExecutorInstance = &DefaultExecutor{
			httpExecutor: NewHTTPExecutor(),
		}
	})
	return defaultExecutorInstance
}

func (e *DefaultExecutor) Execute(ctx context.Context, params interface{}) *model.ExecutionResult {
	// 默认使用HTTP执行器
	return e.httpExecutor.Execute(ctx, params)
}
