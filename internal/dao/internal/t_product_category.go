package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductCategoryDao is the data access object for table t_product_category.
type ProductCategoryDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ProductCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// ProductCategoryColumns defines and stores column names for table t_product_category.
type ProductCategoryColumns struct {
	ID        string // 主键
	OrgID     string // 组织ID
	Name      string // 产品分类名称
	Desc      string // 产品分类描述
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// productCategoryColumns holds the columns for table t_product_category.
var productCategoryColumns = ProductCategoryColumns{
	ID:        "id",
	OrgID:     "org_id",
	Name:      "name",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewProductCategoryDao creates and returns a new DAO object for table data access.
func NewProductCategoryDao() *ProductCategoryDao {
	return &ProductCategoryDao{
		group:   "default",
		table:   "t_product_category",
		columns: productCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductCategoryDao) Columns() ProductCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
