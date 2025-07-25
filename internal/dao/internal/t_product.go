package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductDao is the data access object for table t_product.
type ProductDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ProductColumns // columns contains all the column names of Table for convenient usage.
}

// ProductColumns defines and stores column names for table t_product.
type ProductColumns struct {
	ID         string // 主键
	CategoryID string // 产品分类ID
	OrgID      string // 组织ID
	Name       string // 产品名称
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// productColumns holds the columns for table t_product.
var productColumns = ProductColumns{
	ID:         "id",
	CategoryID: "category_id",
	OrgID:      "org_id",
	Name:       "name",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewProductDao creates and returns a new DAO object for table data access.
func NewProductDao() *ProductDao {
	return &ProductDao{
		group:   "default",
		table:   "t_product",
		columns: productColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductDao) Columns() ProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
