package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetJobLogsReq struct {
	g.Meta `path:"/cron_jobs/logs" tags:"定时任务/日志管理" method:"get" summary:"获取任务执行日志"`
	model.Author
	JobID         string `json:"job_id" dc:"任务ID"`
	ExecuteStatus string `json:"execute_status" v:"in:运行中,运行成功,运行失败#执行状态不正确" dc:"执行状态"`
	model.PageReq
}

type GetJobLogsRes struct {
	List []*CronJobLog `json:"list"`
	model.PageRes
}

type CronJobLog struct {
	ID        int64       `json:"id" dc:"日志ID"`
	JobID     string      `json:"job_id" dc:"任务ID"`
	JobName   string      `json:"job_name" dc:"任务名称"`
	Status    string      `json:"status" dc:"执行状态"`
	Message   string      `json:"message" dc:"执行消息"`
	StartTime *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" dc:"结束时间"`
	Duration  int64       `json:"duration" dc:"执行时长"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
}
