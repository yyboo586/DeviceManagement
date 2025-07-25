package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceDao is the data access object for table t_device.
type DeviceDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns DeviceColumns // columns contains all the column names of Table for convenient usage.
}

// CronJobLogColumns defines and stores column names for table t_cron_job_log.
type DeviceColumns struct {
	ID              string // 主键
	Name            string // 设备名称
	DeviceKey       string // 设备唯一标识
	OrgID           string // 组织ID
	Enabled         string // 设备状态
	OnlineStatus    string // 设备在线状态
	Location        string // 设备位置
	Description     string // 设备描述
	LastOnlineTime  string // 最后在线时间
	LastOfflineTime string // 最后离线时间
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// cronJobLogColumns holds the columns for table t_cron_job_log.
var deviceColumns = DeviceColumns{
	ID:              "id",
	Name:            "name",
	DeviceKey:       "device_key",
	OrgID:           "org_id",
	Enabled:         "enabled",
	OnlineStatus:    "online_status",
	Location:        "location",
	Description:     "description",
	LastOnlineTime:  "last_online_time",
	LastOfflineTime: "last_offline_time",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewDeviceDao creates and returns a new DAO object for table data access.
func NewDeviceDao() *DeviceDao {
	return &DeviceDao{
		group:   "default",
		table:   "t_device",
		columns: deviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceDao) Columns() DeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
