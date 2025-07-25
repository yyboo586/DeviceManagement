package model

import "github.com/gogf/gf/v2/os/gtime"

type ThingModelTemplate struct {
	ID          int64                 `json:"id" dc:"模板标识"`                 // 主键
	OrgID       string                `json:"org_id" dc:"组织ID"`             // 组织ID
	Name        string                `json:"name" dc:"模板名称"`               // 模板名称
	Description string                `json:"description" dc:"模板描述"`        // 模板描述
	Properties  []*ThingModelProperty `json:"properties" dc:"属性定义(JSON格式)"` // 属性定义(JSON格式)
	Services    []*ThingModelService  `json:"services" dc:"服务定义(JSON格式)"`   // 服务定义(JSON格式)
	Events      []*ThingModelEvent    `json:"events" dc:"事件定义(JSON格式)"`     // 事件定义(JSON格式)
	IsSystem    bool                  `json:"is_system" dc:"是否是系统内置模板"`     // 是否是系统内置模板
	CreatedAt   *gtime.Time           `json:"created_at" dc:"创建时间"`         // 创建时间
	UpdatedAt   *gtime.Time           `json:"updated_at" dc:"修改时间"`         // 修改时间
}
