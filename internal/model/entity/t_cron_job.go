package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type TCronJob struct {
	ID                string      `orm:"id"`
	OrgID             string      `orm:"org_id"`
	Name              string      `orm:"name"`
	Params            string      `orm:"params"`
	Remark            string      `orm:"remark"`
	InvokeTarget      string      `orm:"invoke_target"`
	CronExpression    string      `orm:"cron_expression"`
	Enabled           int         `orm:"enabled"`
	LastExecuteStatus int         `orm:"last_execute_status"`
	LastExecuteAt     *gtime.Time `orm:"last_execute_at"`
	NextExecuteAt     *gtime.Time `orm:"next_execute_at"`
	ExecuteCount      int64       `orm:"execute_count"`
	SuccessCount      int64       `orm:"success_count"`
	FailedCount       int64       `orm:"failed_count"`
	CreatedAt         *gtime.Time `orm:"created_at"`
	UpdatedAt         *gtime.Time `orm:"updated_at"`
}
