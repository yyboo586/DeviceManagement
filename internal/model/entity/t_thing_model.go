package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TThingModel 物模型实体
type TThingModel struct {
	ID          int64       `orm:"id"`          // 主键
	ProductID   int64       `orm:"product_id"`  // 产品ID
	OrgID       string      `orm:"org_id"`      // 组织ID
	TemplateID  int64       `orm:"template_id"` // 模板ID
	Name        string      `orm:"name"`        // 物模型名称
	Version     int         `orm:"version"`     // 物模型版本
	Description string      `orm:"description"` // 物模型描述
	Properties  string      `orm:"properties"`  // 属性定义(JSON格式)
	Services    string      `orm:"services"`    // 服务定义(JSON格式)
	Events      string      `orm:"events"`      // 事件定义(JSON格式)
	CreatedAt   *gtime.Time `orm:"created_at"`  // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"`  // 修改时间
}
