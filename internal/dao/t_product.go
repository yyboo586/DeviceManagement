package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// productDao is the data access object for table t_product.
type productDao struct {
	*internal.ProductDao
}

var (
	// Product is globally public accessible object for table t_product operations.
	Product = productDao{
		internal.NewProductDao(),
	}
)
