package model

import "github.com/gogf/gf/v2/os/gtime"

type CronJobTemplate struct {
	ID         string      `json:"id" dc:"模板ID"`
	Name       string      `json:"name" dc:"模板名称"`
	InvokeType string      `json:"invoke_type" dc:"调用类型"`
	Config     string      `json:"config" dc:"配置"`
	CreatedAt  *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updated_at" dc:"更新时间"`
}

var deviceCountJobTemplate = CronJobTemplate{
	Name:       "设备数量统计",
	InvokeType: "http",
	Config:     "{\"url\": \"http://127.0.0.1:9501/api/v1/device-management/devices\"}",
}

var (
	DefaultCronJobList = []CronJobTemplate{
		deviceCountJobTemplate,
	}
)
