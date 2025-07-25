package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ListCronJobTemplateReq struct {
	g.Meta `path:"/cron_jobs/templates" tags:"定时任务/模板管理" method:"get" summary:"列表"`
	model.PageReq
}

type ListCronJobTemplateRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CronJobTemplate `json:"list"`
	model.PageRes
}
