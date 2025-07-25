package model

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

const (
	DeviceCountJobName = "device_count"
)

type CronJobHandler func(ctx context.Context, params interface{}) (success bool, result string)

type CronJobStatus int

const (
	_                     CronJobStatus = iota
	CronJobStatusEnabled                // 启用
	CronJobStatusDisabled               // 禁用
)

type CronJobExecuteStatus int

const (
	CronJobExecuteStatusUnknown CronJobExecuteStatus = iota // 未知状态
	CronJobExecuteStatusRunning                             // 运行中
	CronJobExecuteStatusSuccess                             // 运行成功
	CronJobExecuteStatusFailed                              // 运行失败
)

func GetCronJobExecuteStatusStr(status string) CronJobExecuteStatus {
	switch status {
	case "运行中":
		return CronJobExecuteStatusRunning
	case "运行成功":
		return CronJobExecuteStatusSuccess
	case "运行失败":
		return CronJobExecuteStatusFailed
	default:
		return CronJobExecuteStatusUnknown
	}
}

func GetCronJobExecuteStatus(status CronJobExecuteStatus) string {
	switch status {
	case CronJobExecuteStatusRunning:
		return "运行中"
	case CronJobExecuteStatusSuccess:
		return "运行成功"
	case CronJobExecuteStatusFailed:
		return "运行失败"
	default:
		return "未知状态"
	}
}

type CronJob struct {
	ID                string               `json:"id"`
	OrgID             string               `json:"org_id"`
	TemplateID        string               `json:"template_id"`
	Name              string               `json:"name"`
	Enabled           CronJobStatus        `json:"enabled"`
	Params            string               `json:"params"`
	Remark            string               `json:"remark"`
	InvokeType        string               `json:"invoke_type"`
	CronExpression    string               `json:"cron_expression"`
	Description       string               `json:"description"`
	LastExecuteStatus CronJobExecuteStatus `json:"last_execute_status"`
	LastExecuteAt     *gtime.Time          `json:"last_execute_at"`
	ExecuteCount      int64                `json:"execute_count"`
	SuccessCount      int64                `json:"success_count"`
	FailedCount       int64                `json:"failed_count"`
	CreatedAt         *gtime.Time          `json:"created_at"`
	UpdatedAt         *gtime.Time          `json:"updated_at"`

	Entry *gcron.Entry `json:"-"`
}

type JobStats struct {
	ID                string
	LastExecuteStatus CronJobExecuteStatus `json:"last_execute_status"`
	LastExecuteAt     *gtime.Time          `json:"last_execute_at"`
	ExecuteCount      int64                `json:"execute_count"`
	SuccessCount      int64                `json:"success_count"`
	FailedCount       int64                `json:"failed_count"`
}

type CronJobResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CronJobLog 任务执行日志
type CronJobLog struct {
	ID            int64                `json:"id"`
	OrgID         string               `json:"org_id"`
	JobID         string               `json:"job_id"`
	JobName       string               `json:"job_name"`
	ExecuteStatus CronJobExecuteStatus `json:"execute_status"`
	Result        string               `json:"result"`
	StartTime     *gtime.Time          `json:"start_time"`
	EndTime       *gtime.Time          `json:"end_time"`
	Duration      int64                `json:"duration"` // 执行时长(毫秒)
	CreatedAt     *gtime.Time          `json:"created_at"`
}

type CronJobLogStats struct {
	ID            string               `json:"id"`
	OrgID         string               `json:"org_id"`
	JobID         string               `json:"job_id"`
	ExecuteStatus CronJobExecuteStatus `json:"execute_status"`
	Result        string               `json:"result"`
	EndTime       *gtime.Time          `json:"end_time"`
	Duration      int64                `json:"duration"`
}
