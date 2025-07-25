package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateFromTemplateReq struct {
	g.Meta      `path:"/thing-models/from-template" tags:"物模型管理" method:"post" summary:"从模板创建"`
	OrgID       string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	ProductID   int64  `json:"product_id" v:"required#产品ID不能为空" dc:"产品ID"`
	TemplateID  int64  `json:"template_id" v:"required#模板ID不能为空" dc:"模板ID"`
	Name        string `json:"name" v:"required#物模型名称不能为空" dc:"物模型名称"`
	Description string `json:"description" dc:"物模型描述"`
}

type CreateFromTemplateRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"物模型ID"`
}

type DeleteThingModelReq struct {
	g.Meta `path:"/thing-models/{id}" tags:"物模型管理" method:"delete" summary:"删除"`
	ID     int64 `json:"id" v:"required#物模型ID不能为空" dc:"物模型ID"`
}

type DeleteThingModelRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateThingModelReq struct {
	g.Meta      `path:"/thing-models/{id}" tags:"物模型管理" method:"put" summary:"修改(全量更新)"`
	ID          int64  `json:"id" v:"required#物模型ID不能为空" dc:"物模型ID"`
	ProductID   int64  `json:"product_id" dc:"产品ID"`
	OrgID       string `json:"org_id" dc:"组织ID"`
	TemplateID  int64  `json:"template_id" dc:"模板ID"`
	Name        string `json:"name" dc:"物模型名称"`
	Version     int    `json:"version" dc:"物模型版本"`
	Description string `json:"description" dc:"物模型描述"`
}

type UpdateThingModelRes struct {
	g.Meta `mime:"application/json"`
}

type GetThingModelReq struct {
	g.Meta `path:"/thing-models/{id}" tags:"物模型管理" method:"get" summary:"查询"`
	ID     int64 `json:"id" v:"required#物模型ID不能为空" dc:"物模型ID"`
}

type GetThingModelRes struct {
	g.Meta `mime:"application/json"`
	*model.ThingModel
}

type ListThingModelReq struct {
	g.Meta `path:"/thing-models" tags:"物模型管理" method:"get" summary:"查询物模型列表"`
	OrgID  string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type ListThingModelRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ThingModel `json:"list" dc:"物模型列表"`
	Total  int                 `json:"total" dc:"总数"`
}

type GetThingModelTemplatesReq struct {
	g.Meta `path:"/thing-models/templates" tags:"物模型管理" method:"get" summary:"获取物模型模板列表"`
	OrgID  string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type GetThingModelTemplatesRes struct {
	g.Meta    `mime:"application/json"`
	Templates []*model.ThingModelTemplate `json:"templates" dc:"物模型模板列表"`
}
