package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CronJobDao is the data access object for table t_cron_job.
type CronJobDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CronJobColumns // columns contains all the column names of Table for convenient usage.
}

// CronJobColumns defines and stores column names for table t_cron_job.
type CronJobColumns struct {
	ID                string // 主键
	OrgID             string // 组织ID
	Name              string // 任务名称
	Enabled           string // 是否启用
	Remark            string // 备注
	Params            string // 参数
	InvokeTarget      string // 调用目标
	CronExpression    string // cron执行表达式
	LastExecuteStatus string // 上次执行状态
	LastExecuteAt     string // 上次执行时间
	NextExecuteAt     string // 下次执行时间
	ExecuteCount      string // 执行次数
	SuccessCount      string // 成功次数
	FailedCount       string // 失败次数
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// cronJobColumns holds the columns for table t_cron_job.
var cronJobColumns = CronJobColumns{
	ID:                "id",
	OrgID:             "org_id",
	Name:              "name",
	Enabled:           "enabled",
	Remark:            "remark",
	Params:            "params",
	InvokeTarget:      "invoke_target",
	CronExpression:    "cron_expression",
	LastExecuteStatus: "last_execute_status",
	LastExecuteAt:     "last_execute_at",
	NextExecuteAt:     "next_execute_at",
	ExecuteCount:      "execute_count",
	SuccessCount:      "success_count",
	FailedCount:       "failed_count",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewCronJobDao creates and returns a new DAO object for table data access.
func NewCronJobDao() *CronJobDao {
	return &CronJobDao{
		group:   "default",
		table:   "t_cron_job",
		columns: cronJobColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CronJobDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CronJobDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CronJobDao) Columns() CronJobColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CronJobDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CronJobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CronJobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
