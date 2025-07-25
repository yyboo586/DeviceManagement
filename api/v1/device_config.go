package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddDeviceConfigReq struct {
	g.Meta `path:"/devices/config" tags:"设备管理/配置管理" method:"post" summary:"添加"`
	model.Author
	Type string    `json:"type" v:"required|in:alarm#配置类型不能为空" dc:"配置类型"`
	List []*Config `json:"list" v:"required#配置不能为空" dc:"配置"`
}

type AddDeviceConfigRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteDeviceConfigReq struct {
	g.Meta `path:"/devices/config/{id}" tags:"设备管理/配置管理" method:"delete" summary:"删除"`
	model.Author
	ID int64 `p:"id" v:"required#配置ID不能为空" dc:"配置ID"`
}

type DeleteDeviceConfigRes struct {
	g.Meta `mime:"application/json"`
}

type EditDeviceConfigReq struct {
	g.Meta `path:"/devices/config/{id}" tags:"设备管理/配置管理" method:"put" summary:"编辑"`
	model.Author
	ID     int64   `p:"id" v:"required#配置ID不能为空" dc:"配置ID"`
	Config *Config `json:"config" v:"required#配置不能为空" dc:"配置"`
}

type EditDeviceConfigRes struct {
	g.Meta `mime:"application/json"`
}

type GetDeviceConfigReq struct {
	g.Meta `path:"/devices/config/{id}" tags:"设备管理/配置管理" method:"get" summary:"详情"`
	model.Author
	ID int64 `p:"id" v:"required#配置ID不能为空" dc:"配置ID"`
}

type GetDeviceConfigRes struct {
	g.Meta `mime:"application/json"`
	*DeviceConfig
}

type ListDeviceConfigReq struct {
	g.Meta `path:"/devices/config" tags:"设备管理/配置管理" method:"get" summary:"列表"`
	model.Author
	Type string `json:"type" v:"required|in:alarm#配置类型不能为空" dc:"配置类型"`
}

type ListDeviceConfigRes struct {
	g.Meta `mime:"application/json"`
	List   []*DeviceConfig `json:"list" dc:"配置列表"`
}

type Config struct {
	Key   string `json:"key" v:"required|in:email,phone#配置键不能为空" dc:"配置键"`
	Value string `json:"value" v:"required#配置值不能为空" dc:"配置值"`
}

type DeviceConfig struct {
	ID        int64       `json:"id" dc:"配置ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	Type      string      `json:"type" dc:"配置类型"`
	Key       string      `json:"key" dc:"配置键"`
	Value     string      `json:"value" dc:"配置值"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}
