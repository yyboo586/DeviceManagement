package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// cronJobDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type cronJobDao struct {
	*internal.CronJobDao
}

var (
	// CronJob is globally public accessible object for table tools_gen_table operations.
	CronJob = cronJobDao{
		internal.NewCronJobDao(),
	}
)

// Fill with you ideas below.
