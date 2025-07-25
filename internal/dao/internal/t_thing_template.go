package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThingModelTemplateDao is the data access object for table t_thing_model_template.
type ThingModelTemplateDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ThingModelTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// ThingModelColumns defines and stores column names for table t_thing_model.
type ThingModelTemplateColumns struct {
	ID          string // 主键
	OrgID       string // 组织ID
	Name        string // 物模型名称
	Description string // 物模型描述
	Properties  string // 属性定义(JSON格式)
	Services    string // 服务定义(JSON格式)
	Events      string // 事件定义(JSON格式)
	IsSystem    string // 是否是系统内置模板
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// thingModelColumns holds the columns for table t_thing_model.
var thingModelTemplateColumns = ThingModelTemplateColumns{
	ID:          "id",
	OrgID:       "org_id",
	Name:        "name",
	Description: "description",
	Properties:  "properties",
	Services:    "services",
	Events:      "events",
	IsSystem:    "is_system",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewThingModelDao creates and returns a new DAO object for table data access.
func NewThingModelTemplateDao() *ThingModelTemplateDao {
	return &ThingModelTemplateDao{
		group:   "default",
		table:   "t_thing_model_template",
		columns: thingModelTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ThingModelTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ThingModelTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ThingModelTemplateDao) Columns() ThingModelTemplateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ThingModelTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ThingModelTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ThingModelTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
