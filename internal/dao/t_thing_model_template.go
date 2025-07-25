package dao

import (
	"DeviceManagement/internal/dao/internal"
)

// thingModelTemplateDao is the data access object for table t_thing_model_template.
type thingModelTemplateDao struct {
	*internal.ThingModelTemplateDao
}

var (
	// ThingModelTemplate is globally public accessible object for table t_thing_model_template operations.
	ThingModelTemplate = thingModelTemplateDao{
		internal.NewThingModelTemplateDao(),
	}
)
