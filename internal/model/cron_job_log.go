package model

import "github.com/gogf/gf/v2/os/gtime"

type TCronJobLog struct {
	ID            int64                `json:"id"`
	OrgID         string               `json:"org_id"`
	JobID         string               `json:"job_id"`
	JobName       string               `json:"job_name"`
	ExecuteStatus CronJobExecuteStatus `json:"execute_status"`
	Result        interface{}          `json:"result"`
	StartTime     *gtime.Time          `json:"start_time"`
	EndTime       *gtime.Time          `json:"end_time"`
	Duration      int64                `json:"duration"`
	CreatedAt     *gtime.Time          `json:"created_at"`
}
