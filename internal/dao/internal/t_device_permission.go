package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevicePermissionDao is the data access object for table t_device_permission.
type DevicePermissionDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns DevicePermissionColumns // columns contains all the column names of Table for convenient usage.
}

// CronJobLogColumns defines and stores column names for table t_cron_job_log.
type DevicePermissionColumns struct {
	ID        string // 主键
	OrgID     string // 组织ID
	UserID    string // 用户ID
	DeviceID  string // 设备ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// devicePermissionColumns holds the columns for table t_device_permission.
var devicePermissionColumns = DevicePermissionColumns{
	ID:        "id",
	OrgID:     "org_id",
	UserID:    "user_id",
	DeviceID:  "device_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDeviceDao creates and returns a new DAO object for table data access.
func NewDevicePermissionDao() *DevicePermissionDao {
	return &DevicePermissionDao{
		group:   "default",
		table:   "t_device_permission",
		columns: devicePermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevicePermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevicePermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevicePermissionDao) Columns() DevicePermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevicePermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevicePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevicePermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
