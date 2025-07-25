package dao

import "DeviceManagement/internal/dao/internal"

// devicePermissionDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type devicePermissionDao struct {
	*internal.DevicePermissionDao
}

var (
	DevicePermission = devicePermissionDao{
		internal.NewDevicePermissionDao(),
	}
)

// Fill with you ideas below.
