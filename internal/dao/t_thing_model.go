package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// thingModelDao is the data access object for table t_thing_model.
type thingModelDao struct {
	*internal.ThingModelDao
}

var (
	// ThingModel is globally public accessible object for table t_thing_model operations.
	ThingModel = thingModelDao{
		internal.NewThingModelDao(),
	}
)
