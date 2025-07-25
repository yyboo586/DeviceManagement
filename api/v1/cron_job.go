package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddCronJobReq struct {
	g.Meta         `path:"/cron_job" tags:"定时任务管理" method:"post" summary:"添加定时任务"`
	Name           string `json:"name" v:"required#任务名称不能为空" dc:"任务名称"`
	OrgID          string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	CronExpression string `json:"cron_expression" v:"required#cron表达式不能为空" dc:"cron表达式"`
}

type AddCronJobRes struct {
	ID string `json:"id"`
}

type DeleteCronJobReq struct {
	g.Meta `path:"/cron_job/{id}" tags:"定时任务管理" method:"delete" summary:"删除定时任务"`
	ID     string `p:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

type DeleteCronJobRes struct{}

type EditCronJobReq struct {
	g.Meta         `path:"/cron_job/{id}" tags:"定时任务管理" method:"put" summary:"编辑定时任务"`
	ID             string `p:"id" v:"required#任务ID不能为空" dc:"任务ID"`
	Name           string `json:"name" v:"required#任务名称不能为空" dc:"任务名称"`
	CronExpression string `json:"cron_expression" v:"required#cron表达式不能为空" dc:"cron表达式"`
}

type EditCronJobRes struct{}

type GetCronJobReq struct {
	g.Meta `path:"/cron_job/{id}" tags:"定时任务管理" method:"get" summary:"获取定时任务"`
	ID     string `p:"id"`
}

type GetCronJobRes struct {
	*CronJob
}

type GetCronJobListReq struct {
	g.Meta `path:"/cron_job" tags:"定时任务管理" method:"get" summary:"定时任务列表"`
	OrgID  string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	model.PageReq
}

type GetCronJobListRes struct {
	List []*CronJob `json:"list"`
	model.PageRes
}

// 新增接口
type GetJobLogsReq struct {
	g.Meta `path:"/cron_job/logs" tags:"定时任务管理" method:"get" summary:"获取任务执行日志"`
	OrgID  string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name   string `json:"name" dc:"任务名称"`
	model.PageReq
}

type GetJobLogsRes struct {
	List []*CronJobLog `json:"list"`
	model.PageRes
}

type ExecuteJobReq struct {
	g.Meta `path:"/cron_job/{id}/execute" tags:"定时任务管理" method:"post" summary:"立即执行任务"`
	ID     string `p:"id"`
}

type ExecuteJobRes struct{}

type CronJob struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Params         string      `json:"params"`
	InvokeTarget   string      `json:"invoke_target"`
	CronExpression string      `json:"cron_expression"`
	Enabled        bool        `json:"enabled"`
	LastExecuteAt  *gtime.Time `json:"last_execute_at"`
	NextExecuteAt  *gtime.Time `json:"next_execute_at"`
	ExecuteCount   int64       `json:"execute_count"`
	SuccessCount   int64       `json:"success_count"`
	FailedCount    int64       `json:"failed_count"`
	Remark         string      `json:"remark"`
	CreatedAt      *gtime.Time `json:"created_at"`
	UpdatedAt      *gtime.Time `json:"updated_at"`
}

type CronJobLog struct {
	ID        int64       `json:"id"`
	JobID     string      `json:"job_id"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	StartTime *gtime.Time `json:"start_time"`
	EndTime   *gtime.Time `json:"end_time"`
	Duration  int64       `json:"duration"`
	CreatedAt *gtime.Time `json:"created_at"`
}
