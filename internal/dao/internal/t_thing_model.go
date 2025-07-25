package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThingModelDao is the data access object for table t_thing_model.
type ThingModelDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ThingModelColumns // columns contains all the column names of Table for convenient usage.
}

// ThingModelColumns defines and stores column names for table t_thing_model.
type ThingModelColumns struct {
	ID          string // 主键
	ProductID   string // 产品ID
	OrgID       string // 组织ID
	TemplateID  string // 模板ID
	Name        string // 物模型名称
	Version     string // 物模型版本
	Description string // 物模型描述
	Properties  string // 属性定义(JSON格式)
	Services    string // 服务定义(JSON格式)
	Events      string // 事件定义(JSON格式)
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// thingModelColumns holds the columns for table t_thing_model.
var thingModelColumns = ThingModelColumns{
	ID:          "id",
	ProductID:   "product_id",
	OrgID:       "org_id",
	TemplateID:  "template_id",
	Name:        "name",
	Version:     "version",
	Description: "description",
	Properties:  "properties",
	Services:    "services",
	Events:      "events",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewThingModelDao creates and returns a new DAO object for table data access.
func NewThingModelDao() *ThingModelDao {
	return &ThingModelDao{
		group:   "default",
		table:   "t_thing_model",
		columns: thingModelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ThingModelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ThingModelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ThingModelDao) Columns() ThingModelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ThingModelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ThingModelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ThingModelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
