package logics

import (
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"fmt"
	"sync"
)

// ExecutorType 执行器类型
type ExecutorType string

const (
	ExecutorTypeDefault ExecutorType = "default" // 默认执行器
	ExecutorTypeHTTP    ExecutorType = "http"    // HTTP请求执行器
	ExecutorTypeFunc    ExecutorType = "func"    // 函数执行器
	ExecutorTypeScript  ExecutorType = "script"  // 脚本执行器
)

type ExecutorFactory struct {
	executors map[ExecutorType]service.Executor
	mu        sync.RWMutex
}

var (
	executorFactoryOnce     sync.Once
	executorFactoryInstance *ExecutorFactory
)

func NewExecutorFactory() *ExecutorFactory {
	executorFactoryOnce.Do(func() {
		executorFactoryInstance = &ExecutorFactory{
			executors: make(map[ExecutorType]service.Executor),
		}

		// 注册默认执行器
		executorFactoryInstance.RegisterExecutor(ExecutorTypeHTTP, NewHTTPExecutor())
		executorFactoryInstance.RegisterExecutor(ExecutorTypeDefault, NewDefaultExecutor())
	})
	return executorFactoryInstance
}

// RegisterExecutor 注册执行器
func (f *ExecutorFactory) RegisterExecutor(executorType ExecutorType, executor service.Executor) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.executors[executorType] = executor
}

// GetExecutor 获取执行器
func (f *ExecutorFactory) GetExecutor(executorType ExecutorType) (service.Executor, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	executor, exists := f.executors[executorType]
	if !exists {
		return nil, fmt.Errorf("executor type %s not found", executorType)
	}

	return executor, nil
}

// Execute 执行任务
func (f *ExecutorFactory) Execute(ctx context.Context, executorType ExecutorType, params interface{}) *model.ExecutionResult {
	executor, err := f.GetExecutor(executorType)
	if err != nil {
		return &model.ExecutionResult{
			Success: false,
			Message: fmt.Sprintf("获取执行器失败: %v", err),
			Error:   err,
		}
	}

	return executor.Execute(ctx, params)
}

// ScriptExecutor 脚本执行器
type ScriptExecutor struct {
	// 可以添加脚本执行相关的依赖
}

func NewScriptExecutor() *ScriptExecutor {
	return &ScriptExecutor{}
}

func (e *ScriptExecutor) Execute(ctx context.Context, params interface{}) *model.ExecutionResult {
	// 实现脚本执行逻辑
	// 这里可以执行shell脚本、Python脚本等
	return &model.ExecutionResult{
		Success: true,
		Message: "脚本执行成功",
		Data:    "脚本执行结果",
	}
}

// FunctionExecutor 函数执行器
type FunctionExecutor struct {
	// 可以添加函数执行相关的依赖
}

func NewFunctionExecutor() *FunctionExecutor {
	return &FunctionExecutor{}
}

func (e *FunctionExecutor) Execute(ctx context.Context, params interface{}) *model.ExecutionResult {
	// 实现函数执行逻辑
	// 这里可以执行Go函数、远程函数调用等
	return &model.ExecutionResult{
		Success: true,
		Message: "函数执行成功",
		Data:    "函数执行结果",
	}
}
