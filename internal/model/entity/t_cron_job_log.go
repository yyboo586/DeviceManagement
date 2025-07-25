package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type TCronJobLog struct {
	ID            int64       `orm:"id"`
	OrgID         string      `orm:"org_id"`
	JobID         string      `orm:"job_id"`
	ExecuteStatus int         `orm:"execute_status"`
	Result        string      `orm:"result"`
	StartTime     *gtime.Time `orm:"start_time"`
	EndTime       *gtime.Time `orm:"end_time"`
	Duration      int64       `orm:"duration"`
	CreatedAt     *gtime.Time `orm:"created_at"`
}
