package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// sysConfigDao is the data access object for table sys_config.
// You can define custom methods on it to extend its functionality as you wish.
type productCategoryDao struct {
	*internal.ProductCategoryDao
}

var (
	// SysConfig is globally public accessible object for table sys_config operations.
	ProductCategory = productCategoryDao{
		internal.NewProductCategoryDao(),
	}
)

// Fill with you ideas below.
