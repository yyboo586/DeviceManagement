package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// deviceConfigDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type deviceConfigDao struct {
	*internal.DeviceConfigDao
}

var (
	// DeviceConfig is globally public accessible object for table tools_gen_table operations.
	DeviceConfig = deviceConfigDao{
		internal.NewDeviceConfigDao(),
	}
)

// Fill with you ideas below.
