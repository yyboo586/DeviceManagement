package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// cronJobTemplateDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type cronJobTemplateDao struct {
	*internal.CronJobTemplateDao
}

var (
	// CronJobTemplate is globally public accessible object for table t_cron_job_template operations.
	CronJobTemplate = cronJobTemplateDao{
		internal.NewCronJobTemplateDao(),
	}
)

// Fill with you ideas below.
