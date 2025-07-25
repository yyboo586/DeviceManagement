package model

import (
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

type CronJobStatus int

const (
	_                     CronJobStatus = iota
	CronJobStatusEnabled                // 启用
	CronJobStatusDisabled               // 禁用
)

type CronJobExecuteStatus int

const (
	_                           CronJobExecuteStatus = iota // 等待执行
	CronJobExecuteStatusRunning                             // 运行中
	CronJobExecuteStatusSuccess                             // 执行成功
	CronJobExecuteStatusFailed                              // 执行失败
)

type CronJob struct {
	ID                string               `json:"id"`
	OrgID             string               `json:"org_id"`
	Name              string               `json:"name"`
	Enabled           CronJobStatus        `json:"enabled"`
	Params            string               `json:"params"`
	Remark            string               `json:"remark"`
	InvokeTarget      string               `json:"invoke_target"`
	CronExpression    string               `json:"cron_expression"`
	LastExecuteStatus CronJobExecuteStatus `json:"last_execute_status"`
	LastExecuteAt     *gtime.Time          `json:"last_execute_at"`
	NextExecuteAt     *gtime.Time          `json:"next_execute_at"`
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
	NextExecuteAt     *gtime.Time          `json:"next_execute_at"`
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
	ExecuteStatus CronJobExecuteStatus `json:"execute_status"`
	Result        interface{}          `json:"result"`
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
	Result        interface{}          `json:"result"`
	EndTime       *gtime.Time          `json:"end_time"`
	Duration      int64                `json:"duration"`
}
