package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceConfigDao is the data access object for table t_device_config.
type DeviceConfigDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns DeviceConfigColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceConfigColumns defines and stores column names for table t_device_config.
type DeviceConfigColumns struct {
	ID        string // 配置ID
	OrgID     string // 组织ID
	Type      string // 配置类型
	Key       string // 配置键
	Value     string // 配置值
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// deviceConfigColumns holds the columns for table t_device_config.
var deviceConfigColumns = DeviceConfigColumns{
	ID:        "id",
	OrgID:     "org_id",
	Type:      "type",
	Key:       "key",
	Value:     "value",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCronJobDao creates and returns a new DAO object for table data access.
func NewDeviceConfigDao() *DeviceConfigDao {
	return &DeviceConfigDao{
		group:   "default",
		table:   "t_device_config",
		columns: deviceConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceConfigDao) Columns() DeviceConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
