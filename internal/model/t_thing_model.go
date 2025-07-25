package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ThingModel struct {
	ID          int64                 `json:"id"`          // 物模型ID
	ProductID   int64                 `json:"product_id"`  // 产品ID
	OrgID       string                `json:"org_id"`      // 组织ID
	TemplateID  int64                 `json:"template_id"` // 模板ID
	Name        string                `json:"name"`        // 物模型名称
	Version     int                   `json:"version"`     // 物模型版本
	Description string                `json:"description"` // 物模型描述
	Properties  []*ThingModelProperty `json:"properties"`  // 属性定义(JSON格式)
	Services    []*ThingModelService  `json:"services"`    // 服务定义(JSON格式)
	Events      []*ThingModelEvent    `json:"events"`      // 事件定义(JSON格式)
	CreatedAt   *gtime.Time           `json:"created_at"`  // 创建时间
	UpdatedAt   *gtime.Time           `json:"updated_at"`  // 修改时间
}

type ThingModelProperty struct {
	ID          string  `json:"id" dc:"属性标识"`
	Name        string  `json:"name" dc:"属性名称"`
	DataType    string  `json:"data_type" dc:"数据类型(int,float,string,bool,enum)"`
	Unit        string  `json:"unit" dc:"单位"`
	Min         float64 `json:"min" dc:"最小值"`
	Max         float64 `json:"max" dc:"最大值"`
	Required    bool    `json:"required" dc:"是否必填"`
	Description string  `json:"description" dc:"描述"`
}

type ThingModelService struct {
	ID          string                `json:"id" dc:"服务标识"`
	Name        string                `json:"name" dc:"服务名称"`
	Description string                `json:"description" dc:"描述"`
	Input       []*ThingModelProperty `json:"input" dc:"输入参数"`
	Output      []*ThingModelProperty `json:"output" dc:"输出参数"`
}

type ThingModelEvent struct {
	ID          string                `json:"id" dc:"事件标识"`
	Name        string                `json:"name" dc:"事件名称"`
	Description string                `json:"description" dc:"事件描述"`
	Output      []*ThingModelProperty `json:"output" dc:"输出参数"`
}
