package dao

import "DeviceManagement/internal/dao/internal"

// deviceDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type deviceDao struct {
	*internal.DeviceDao
}

var (
	Device = deviceDao{
		internal.NewDeviceDao(),
	}
)

// Fill with you ideas below.
