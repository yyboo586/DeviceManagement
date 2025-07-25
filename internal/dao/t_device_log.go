package dao

import "DeviceManagement/internal/dao/internal"

// deviceLogDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type deviceLogDao struct {
	*internal.DeviceLogDao
}

var (
	// DeviceLog is globally public accessible object for table tools_gen_table operations.
	DeviceLog = deviceLogDao{
		internal.NewDeviceLogDao(),
	}
)

// Fill with you ideas below.
