package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	devicePermissionOnce     sync.Once
	devicePermissionInstance *devicePermissionLogic
)

type devicePermissionLogic struct {
}

func NewDevicePermission() service.IDevicePermissionService {
	devicePermissionOnce.Do(func() {
		devicePermissionInstance = &devicePermissionLogic{}
	})
	return devicePermissionInstance
}

func (l *devicePermissionLogic) BindDevicePermission(ctx context.Context, req *v1.BindDevicePermissionReq) (err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	var dataInserts g.Slice
	for _, deviceID := range req.DeviceIDs {
		dataInserts = append(dataInserts, map[string]interface{}{
			dao.DevicePermission.Columns().OrgID:    operatorInfo.OrgID,
			dao.DevicePermission.Columns().UserID:   req.UserID,
			dao.DevicePermission.Columns().DeviceID: deviceID,
		})
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除原有的权限
		_, err = dao.DevicePermission.Ctx(ctx).Where(dao.DevicePermission.Columns().OrgID, operatorInfo.OrgID).Where(dao.DevicePermission.Columns().UserID, req.UserID).Delete()
		if err != nil {
			return err
		}
		// 添加新的权限
		_, err = dao.DevicePermission.Ctx(ctx).Data(dataInserts).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

func (l *devicePermissionLogic) CheckPermission(ctx context.Context, deviceInfo *model.Device) (ok bool, err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	// 检查当前用户的角色
	for _, roleMap := range operatorInfo.Roles {
		for _, roleName := range roleMap {
			if roleName == "SuperAdmin" || roleName == "FrontOrgAdmin" {
				ok = true
				return
			}
		}
	}

	// 检查当前用户是否是设备创建者
	if deviceInfo.CreatorID == operatorInfo.UserID {
		ok = true
		return
	}

	// 检查当前用户的设备权限
	exist, err := dao.DevicePermission.Ctx(ctx).Where(dao.DevicePermission.Columns().UserID, operatorInfo.UserID).Where(dao.DevicePermission.Columns().DeviceID, deviceInfo.ID).Exist()
	if err != nil {
		return false, err
	}
	if exist {
		ok = true
	}

	return
}
