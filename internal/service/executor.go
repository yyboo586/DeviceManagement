package service

import (
	"DeviceManagement/internal/model"
	"context"
)

// Executor 任务执行器接口
type Executor interface {
	Execute(ctx context.Context, params interface{}) *model.ExecutionResult
}
