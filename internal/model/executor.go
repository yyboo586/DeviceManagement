package model

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// ExecutionResult 执行结果
type ExecutionResult struct {
	Success bool        `json:"success"` // 是否成功
	Message string      `json:"message"` // 执行消息
	Data    interface{} `json:"data"`    // 执行返回的数据
	Error   error       `json:"error"`   // 错误信息

	Duration  int64       `json:"duration"`  // 执行时长(毫秒)
	StartTime *gtime.Time `json:"startTime"` // 开始时间
	EndTime   *gtime.Time `json:"endTime"`   // 结束时间
}

// HTTPRequestConfig HTTP请求配置
type HTTPRequestConfig struct {
	Method  string                 `json:"method"`  // HTTP方法: GET, POST, PUT, DELETE等
	URL     string                 `json:"url"`     // 请求地址
	Headers map[string]interface{} `json:"headers"` // 请求头
	Body    interface{}            `json:"body"`    // 请求体
	Timeout time.Duration          `json:"timeout"` // 超时时间
	Params  map[string]interface{} `json:"params"`  // 额外参数
}
