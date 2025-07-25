package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TimeUnitConfig 用于描述Cron表达式中单个时间单位（如分钟、小时、日期等）的配置
type TimeUnitConfig struct {
	// Type 表达式类型，决定了其他字段的含义
	// all: 全匹配（*），表示该时间单位的所有值
	// specific: 特定值（如1,3,5），用逗号分隔的离散值
	// range: 连续范围（如1-5），从Start到End的闭区间
	// interval: 间隔步长（如*/5或1-10/2），在范围内按Step间隔执行
	Type string `json:"type" v:"in:all,specific,range,interval#类型必须为all/specific/range/interval" dc:"时间单位的表达式类型"`

	// SpecificValues 特定值列表，仅Type=specific时有效
	// 示例：[1,3,5] 对应Cron中的"1,3,5"
	SpecificValues []int `json:"specific_values" dc:"离散的特定值列表"`

	// Start 起始值，Type=range/interval时有效
	// 示例：range类型中1-5的Start=1；interval类型中1-10/2的Start=1
	Start int `json:"start" dc:"范围或间隔的起始值"`

	// End 结束值，Type=range/interval时有效（interval类型可选，无则表示到最大值）
	// 示例：range类型中1-5的End=5；interval类型中1-10/2的End=10
	End int `json:"end" dc:"范围或间隔的结束值"`

	// Step 步长，Type=interval时有效（必须>0）
	// 示例：*/5的Step=5；1-10/2的Step=2
	Step int `json:"step" dc:"间隔步长(>0)"`
}

type TaskTimeConfig struct {
	Second *TimeUnitConfig `json:"second" v:"required#秒不能为空" dc:"秒"`
	Minute *TimeUnitConfig `json:"minute" v:"required#分钟不能为空" dc:"分钟"`
	Hour   *TimeUnitConfig `json:"hour" v:"required#小时不能为空" dc:"小时"`
	Day    *TimeUnitConfig `json:"day" v:"required#日不能为空" dc:"日"`
	Week   *TimeUnitConfig `json:"week" v:"required#周不能为空" dc:"周"`
}

type AddCronJobReq struct {
	g.Meta `path:"/cron_jobs" tags:"定时任务" method:"post" summary:"添加"`
	model.Author
	TemplateID string          `json:"template_id" v:"required#任务模板ID不能为空" dc:"任务模板ID"`
	OrgID      string          `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Remark     string          `json:"remark" dc:"任务备注"`
	TimeConfig *TaskTimeConfig `json:"time_config" v:"required#时间配置不能为空" dc:"时间配置"`
}

type AddCronJobRes struct {
	g.Meta         `mime:"application/json"`
	ID             string `json:"id" dc:"任务ID"`
	CronExpression string `json:"cron_expression" dc:"cron表达式"`
	Description    string `json:"description" dc:"cron描述"`
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
	ID         string          `p:"id" v:"required#任务ID不能为空" dc:"任务ID"`
	Name       string          `json:"name" v:"required#任务名称不能为空" dc:"任务名称"`
	Remark     string          `json:"remark" dc:"任务备注"`
	TimeConfig *TaskTimeConfig `json:"time_config" v:"required#时间配置不能为空" dc:"时间配置"`
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
	InvokeType     string      `json:"invoke_type" dc:"调用类型"`
	CronExpression string      `json:"cron_expression" dc:"cron表达式"`
	Description    string      `json:"description" dc:"cron描述"`
	Enabled        bool        `json:"enabled" dc:"是否启用"`
	LastExecuteAt  *gtime.Time `json:"last_execute_at" dc:"上次执行时间"`
	ExecuteCount   int64       `json:"execute_count" dc:"执行次数"`
	SuccessCount   int64       `json:"success_count" dc:"成功次数"`
	FailedCount    int64       `json:"failed_count" dc:"失败次数"`
	Remark         string      `json:"remark" dc:"备注"`
	CreatedAt      *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updated_at" dc:"更新时间"`
}
