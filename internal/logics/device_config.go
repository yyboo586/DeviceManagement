package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"DeviceManagement/internal/service"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	deviceConfigOnce          sync.Once
	deviceConfigLogicInstance *deviceConfigLogic
)

type deviceConfigLogic struct{}

func NewDeviceConfig() service.IDeviceConfig {
	deviceConfigOnce.Do(func() {
		deviceConfigLogicInstance = &deviceConfigLogic{}
	})
	return deviceConfigLogicInstance
}

func (l *deviceConfigLogic) Add(ctx context.Context, in *v1.AddDeviceConfigReq) (err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	for _, config := range in.List {
		typ := model.GetDeviceConfigType(in.Type)
		if typ == model.DeviceConfigTypeUnknown {
			err = fmt.Errorf("不支持的配置类型")
			g.Log().Error(ctx, err)
			return
		}

		insertData := g.Map{
			dao.DeviceConfig.Columns().OrgID: operatorInfo.OrgID,
			dao.DeviceConfig.Columns().Type:  typ,
			dao.DeviceConfig.Columns().Key:   config.Key,
			dao.DeviceConfig.Columns().Value: config.Value,
		}
		_, err = dao.DeviceConfig.Ctx(ctx).Data(insertData).Insert()
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = fmt.Errorf("%s配置已存在", config.Key)
			}
			g.Log().Error(ctx, err)
			return
		}
	}

	return
}

func (l *deviceConfigLogic) Delete(ctx context.Context, id int64) (err error) {
	_, err = dao.DeviceConfig.Ctx(ctx).Where(dao.DeviceConfig.Columns().ID, id).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (l *deviceConfigLogic) Edit(ctx context.Context, in *v1.EditDeviceConfigReq) (err error) {
	dataUpdate := g.Map{
		dao.DeviceConfig.Columns().Key:   in.Config.Key,
		dao.DeviceConfig.Columns().Value: in.Config.Value,
	}
	_, err = dao.DeviceConfig.Ctx(ctx).Where(dao.DeviceConfig.Columns().ID, in.ID).Data(dataUpdate).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (l *deviceConfigLogic) Get(ctx context.Context, id int64) (out *model.DeviceConfig, err error) {
	var entity entity.TDeviceConfig
	err = dao.DeviceConfig.Ctx(ctx).Where(dao.DeviceConfig.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("配置不存在")
		}
		g.Log().Error(ctx, err)
		return
	}
	return l.convertEntityToLogics(&entity), nil
}

func (l *deviceConfigLogic) List(ctx context.Context, in *v1.ListDeviceConfigReq) (out []*model.DeviceConfig, err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	typ := model.GetDeviceConfigType(in.Type)
	if typ == model.DeviceConfigTypeUnknown {
		err = fmt.Errorf("配置类型不存在")
		g.Log().Error(ctx, err)
		return
	}

	var entities []*entity.TDeviceConfig
	err = dao.DeviceConfig.Ctx(ctx).Where(dao.DeviceConfig.Columns().OrgID, operatorInfo.OrgID).Where(dao.DeviceConfig.Columns().Type, typ).Scan(&entities)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		g.Log().Error(ctx, err)
		return
	}

	for _, entity := range entities {
		out = append(out, l.convertEntityToLogics(entity))
	}
	return
}

func (l *deviceConfigLogic) convertEntityToLogics(in *entity.TDeviceConfig) (out *model.DeviceConfig) {
	out = &model.DeviceConfig{
		ID:        in.ID,
		OrgID:     in.OrgID,
		Type:      model.DeviceConfigType(in.Type),
		Key:       in.Key,
		Value:     in.Value,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
