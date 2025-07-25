package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CronJobTemplateDao is the data access object for table t_cron_job_template.
type CronJobTemplateDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns CronJobTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// CronJobTemplateColumns defines and stores column names for table t_cron_job_template.
type CronJobTemplateColumns struct {
	ID         string // 主键
	Name       string // 任务名称
	InvokeType string // 调用类型
	Config     string // 配置
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// cronJobTemplateColumns holds the columns for table t_cron_job_template.
var cronJobTemplateColumns = CronJobTemplateColumns{
	ID:         "id",
	Name:       "name",
	InvokeType: "invoke_type",
	Config:     "config",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewCronJobDao creates and returns a new DAO object for table data access.
func NewCronJobTemplateDao() *CronJobTemplateDao {
	return &CronJobTemplateDao{
		group:   "default",
		table:   "t_cron_job_template",
		columns: cronJobTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CronJobTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CronJobTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CronJobTemplateDao) Columns() CronJobTemplateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CronJobTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CronJobTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CronJobTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
