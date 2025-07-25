package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceOnlineLogDao is the data access object for table t_device_online_log.
type DeviceOnlineLogDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns DeviceOnlineLogColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceOnlineLogColumns defines and stores column names for table t_device_online_log.
type DeviceOnlineLogColumns struct {
	ID           string // 主键
	DeviceID     string // 设备ID
	DeviceKey    string // 设备唯一标识
	OrgID        string // 组织ID
	EventType    string // 事件类型
	OnlineStatus string // 设备在线状态
	IPAddress    string // 设备IP地址
	ClientID     string // 客户端ID
	Reason       string // 上下线原因
	Duration     string // 在线时长
	CreatedAt    string // 创建时间
}

// deviceOnlineLogColumns holds the columns for table t_device_online_log.
var deviceOnlineLogColumns = DeviceOnlineLogColumns{
	ID:           "id",
	DeviceID:     "device_id",
	DeviceKey:    "device_key",
	OrgID:        "org_id",
	EventType:    "event_type",
	OnlineStatus: "online_status",
	IPAddress:    "ip_address",
	ClientID:     "client_id",
	Reason:       "reason",
	Duration:     "duration",
	CreatedAt:    "created_at",
}

// NewDeviceOnlineLogDao creates and returns a new DAO object for table data access.
func NewDeviceOnlineLogDao() *DeviceOnlineLogDao {
	return &DeviceOnlineLogDao{
		group:   "default",
		table:   "t_device_online_log",
		columns: deviceOnlineLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceOnlineLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceOnlineLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceOnlineLogDao) Columns() DeviceOnlineLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceOnlineLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceOnlineLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if the function f returns error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceOnlineLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
