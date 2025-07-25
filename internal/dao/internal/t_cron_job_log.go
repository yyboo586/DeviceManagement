package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CronJobLogDao is the data access object for table t_cron_job_log.
type CronJobLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CronJobLogColumns // columns contains all the column names of Table for convenient usage.
}

// CronJobLogColumns defines and stores column names for table t_cron_job_log.
type CronJobLogColumns struct {
	ID            string // 主键
	OrgID         string // 组织ID
	JobID         string // 任务ID
	JobName       string // 任务名称
	ExecuteStatus string // 执行状态
	Result        string // 执行消息
	StartTime     string // 开始时间
	EndTime       string // 结束时间
	Duration      string // 执行时长
	CreatedAt     string // 创建时间
}

// cronJobLogColumns holds the columns for table t_cron_job_log.
var cronJobLogColumns = CronJobLogColumns{
	ID:            "id",
	OrgID:         "org_id",
	JobID:         "job_id",
	JobName:       "job_name",
	ExecuteStatus: "execute_status",
	Result:        "result",
	StartTime:     "start_time",
	EndTime:       "end_time",
	Duration:      "duration",
	CreatedAt:     "created_at",
}

// NewCronJobLogDao creates and returns a new DAO object for table data access.
func NewCronJobLogDao() *CronJobLogDao {
	return &CronJobLogDao{
		group:   "default",
		table:   "t_cron_job_log",
		columns: cronJobLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CronJobLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CronJobLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CronJobLogDao) Columns() CronJobLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CronJobLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CronJobLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CronJobLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
