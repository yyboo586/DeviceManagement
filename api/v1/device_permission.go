package v1

import (
	"DeviceManagement/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type BindDevicePermissionReq struct {
	g.Meta `path:"/devices/permissions" tags:"设备管理/权限管理" method:"post" summary:"绑定权限"`
	model.Author
	UserID    string  `json:"user_id" v:"required#用户ID不能为空" dc:"用户ID"`
	DeviceIDs []int64 `json:"device_ids" v:"required#设备ID不能为空" dc:"设备ID"`
}

type BindDevicePermissionRes struct {
	g.Meta `mime:"application/json"`
}
