package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// cronJobLogDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type cronJobLogDao struct {
	*internal.CronJobLogDao
}

var (
	// CronJobLog is globally public accessible object for table t_cron_job_log operations.
	CronJobLog = cronJobLogDao{
		internal.NewCronJobLogDao(),
	}
)

// Fill with you ideas below.
