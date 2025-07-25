package dao

import "DeviceManagement/internal/dao/internal"

// deviceDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type deviceOnlineLogDao struct {
	*internal.DeviceOnlineLogDao
}

var (
	// CronJob is globally public accessible object for table tools_gen_table operations.
	DeviceOnlineLog = deviceOnlineLogDao{
		internal.NewDeviceOnlineLogDao(),
	}
)

// Fill with you ideas below.
