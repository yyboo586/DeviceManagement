package entity

import "github.com/gogf/gf/v2/os/gtime"

type TThingModelTemplate struct {
	ID          int64       `orm:"id"`          // 主键
	OrgID       string      `orm:"org_id"`      // 组织ID
	Name        string      `orm:"name"`        // 物模型名称
	Description string      `orm:"description"` // 物模型描述
	Properties  string      `orm:"properties"`  // 属性定义(JSON格式)
	Services    string      `orm:"services"`    // 服务定义(JSON格式)
	Events      string      `orm:"events"`      // 事件定义(JSON格式)
	IsSystem    int         `orm:"is_system"`   // 是否是系统内置模板
	CreatedAt   *gtime.Time `orm:"created_at"`  // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"`  // 修改时间
}
