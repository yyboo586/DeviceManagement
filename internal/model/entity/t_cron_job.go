package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type TCronJob struct {
	ID                string      `orm:"id"`
	OrgID             string      `orm:"org_id"`
	TemplateID        string      `orm:"template_id"`
	Name              string      `orm:"name"`
	Params            string      `orm:"params"`
	Remark            string      `orm:"remark"`
	InvokeType        string      `orm:"invoke_type"`
	CronExpression    string      `orm:"cron_expression"`
	Description       string      `orm:"description"`
	Enabled           int         `orm:"enabled"`
	LastExecuteStatus int         `orm:"last_execute_status"`
	LastExecuteAt     *gtime.Time `orm:"last_execute_at"`
	ExecuteCount      int64       `orm:"execute_count"`
	SuccessCount      int64       `orm:"success_count"`
	FailedCount       int64       `orm:"failed_count"`
	CreatedAt         *gtime.Time `orm:"created_at"`
	UpdatedAt         *gtime.Time `orm:"updated_at"`
}
