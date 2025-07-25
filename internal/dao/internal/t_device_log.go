package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceLogDao is the data access object for table t_device_log.
type DeviceLogDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns DeviceLogColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceLogColumns defines and stores column names for table t_device_log.
type DeviceLogColumns struct {
	ID         string // 主键
	OrgID      string // 组织ID
	DeviceID   string // 设备ID
	DeviceName string // 设备名称
	DeviceKey  string // 设备唯一标识
	Type       string // 日志类型
	Content    string // 日志内容
	Timestamp  string // 时间戳
	CreatedAt  string // 创建时间
}

// deviceLogColumns holds the columns for table t_device_log.
var deviceLogColumns = DeviceLogColumns{
	ID:         "id",
	OrgID:      "org_id",
	DeviceID:   "device_id",
	DeviceName: "device_name",
	DeviceKey:  "device_key",
	Type:       "type",
	Content:    "content",
	Timestamp:  "timestamp",
	CreatedAt:  "created_at",
}

// NewDeviceLogDao creates and returns a new DAO object for table data access.
func NewDeviceLogDao() *DeviceLogDao {
	return &DeviceLogDao{
		group:   "default",
		table:   "t_device_log",
		columns: deviceLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceLogDao) Columns() DeviceLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if the function f returns error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
