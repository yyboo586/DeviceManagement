package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddCronJobReq struct {
	g.Meta `path:"/cron_jobs" tags:"定时任务" method:"post" summary:"添加"`
	model.Author
	Name           string `json:"name" v:"required#任务名称不能为空" dc:"任务名称"`
	OrgID          string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	CronExpression string `json:"cron_expression" v:"required#cron表达式不能为空" dc:"cron表达式"`
}

type AddCronJobRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id"`
}

type DeleteCronJobReq struct {
	g.Meta `path:"/cron_jobs/{id}" tags:"定时任务" method:"delete" summary:"删除"`
	model.Author
	ID string `p:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

type DeleteCronJobRes struct {
	g.Meta `mime:"application/json"`
}

type EditCronJobReq struct {
	g.Meta `path:"/cron_jobs/{id}" tags:"定时任务" method:"put" summary:"编辑"`
	model.Author
	ID             string `p:"id" v:"required#任务ID不能为空" dc:"任务ID"`
	Name           string `json:"name" v:"required#任务名称不能为空" dc:"任务名称"`
	CronExpression string `json:"cron_expression" v:"required#cron表达式不能为空" dc:"cron表达式"`
}

type EditCronJobRes struct {
	g.Meta `mime:"application/json"`
}

type GetCronJobReq struct {
	g.Meta `path:"/cron_jobs/{id}" tags:"定时任务" method:"get" summary:"详情"`
	model.Author
	ID string `p:"id"`
}

type GetCronJobRes struct {
	g.Meta `mime:"application/json"`
	*CronJob
}

type GetCronJobListReq struct {
	g.Meta `path:"/cron_jobs" tags:"定时任务" method:"get" summary:"列表"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	model.PageReq
}

type GetCronJobListRes struct {
	g.Meta `mime:"application/json"`
	List   []*CronJob `json:"list"`
	model.PageRes
}

type ExecuteJobReq struct {
	g.Meta `path:"/cron_jobs/{job_id}/execute" tags:"定时任务" method:"post" summary:"立即执行任务(异步获取执行结果)"`
	model.Author
	JobID string `p:"job_id"`
}

type ExecuteJobRes struct {
	g.Meta `mime:"application/json"`
}

type CronJob struct {
	ID             string      `json:"id" dc:"任务ID"`
	Name           string      `json:"name" dc:"任务名称"`
	Params         string      `json:"params" dc:"参数"`
	InvokeTarget   string      `json:"invoke_target" dc:"调用目标"`
	CronExpression string      `json:"cron_expression" dc:"cron表达式"`
	Enabled        bool        `json:"enabled" dc:"是否启用"`
	LastExecuteAt  *gtime.Time `json:"last_execute_at" dc:"上次执行时间"`
	NextExecuteAt  *gtime.Time `json:"next_execute_at" dc:"下次执行时间"`
	ExecuteCount   int64       `json:"execute_count" dc:"执行次数"`
	SuccessCount   int64       `json:"success_count" dc:"成功次数"`
	FailedCount    int64       `json:"failed_count" dc:"失败次数"`
	Remark         string      `json:"remark" dc:"备注"`
	CreatedAt      *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updated_at" dc:"更新时间"`
}
