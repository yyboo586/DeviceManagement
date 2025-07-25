package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
)

type device struct {
	pool *grpool.Pool
}

func NewDevice() *device {
	return &device{
		pool: grpool.New(100),
	}
}

func (d *device) InvokeAlarm(ctx context.Context, orgID string, content string) (err error) {
	d.pool.Add(ctx, func(ctx context.Context) {
		d.Alarm(ctx, orgID, content)
	})
	return
}

func (d *device) Add(ctx context.Context, in *model.Device) (id int64, err error) {
	insertData := map[string]interface{}{
		dao.Device.Columns().ProductID:   in.ProductID,
		dao.Device.Columns().OrgID:       in.OrgID,
		dao.Device.Columns().CreatorID:   in.CreatorID,
		dao.Device.Columns().DeviceKey:   in.DeviceKey,
		dao.Device.Columns().Name:        in.Name,
		dao.Device.Columns().Enabled:     model.DeviceStatusEnabled,
		dao.Device.Columns().Location:    in.Location,
		dao.Device.Columns().Description: in.Description,
	}

	id, err = dao.Device.Ctx(ctx).Data(insertData).InsertAndGetId()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("设备已存在")
		}
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (d *device) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.Device.Ctx(ctx).WhereIn(dao.Device.Columns().ID, ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (d *device) Update(ctx context.Context, in *model.Device) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().ProductID:   in.ProductID,
		dao.Device.Columns().Name:        in.Name,
		dao.Device.Columns().Location:    in.Location,
		dao.Device.Columns().Description: in.Description,
	}

	_, err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().ID, in.ID).Data(updateData).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (d *device) EditDeviceStatus(ctx context.Context, ids []int64, enabled bool) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().Enabled: model.DeviceStatusEnabled,
	}
	if !enabled {
		updateData[dao.Device.Columns().Enabled] = model.DeviceStatusDisabled
	}

	_, err = dao.Device.Ctx(ctx).WhereIn(dao.Device.Columns().ID, ids).Data(updateData).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (d *device) Get(ctx context.Context, id int64) (out *model.Device, err error) {
	var device entity.TDevice
	err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().ID, id).Scan(&device)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("设备不存在")
		}
		g.Log().Error(ctx, err)
		return
	}

	out = d.convertDeviceToLogic(&device)
	ok, err := devicePermissionInstance.CheckPermission(ctx, out)
	if err != nil {
		out = nil
		g.Log().Error(ctx, err)
		return
	}
	if !ok {
		out = nil
		err = fmt.Errorf("无权限访问该设备")
		g.Log().Error(ctx, err)
		return
	}

	return out, nil
}

func (d *device) List(ctx context.Context, orgID string, page *model.PageReq) (out []*model.Device, pageRes *model.PageRes, err error) {
	if page.Page == 0 {
		page.Page = 1
	}
	if page.PageSize == 0 {
		page.PageSize = model.DefaultPageSize
	}

	m := dao.Device.Ctx(ctx).Where(dao.Device.Columns().OrgID, orgID)

	// 获取总数
	total, err := m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 获取列表
	var devices []*entity.TDevice
	err = m.Page(page.Page, page.PageSize).Scan(&devices)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, device := range devices {
		item := d.convertDeviceToLogic(device)
		ok, err := devicePermissionInstance.CheckPermission(ctx, item)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if !ok {
			continue
		}
		out = append(out, item)
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: page.Page,
	}
	return
}

func (d *device) Alarm(ctx context.Context, orgID string, content string) (err error) {
	configs, err := deviceConfigLogicInstance.List(ctx, &v1.ListDeviceConfigReq{
		Type: "alarm",
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if len(configs) == 0 {
		err = fmt.Errorf("配置不存在")
		g.Log().Error(ctx, err)
		return
	}

	var config *model.DeviceConfig
	for _, v := range configs {
		if v.Type == model.DeviceConfigTypeAlarm && v.Key == "email" {
			config = v
			break
		}
	}

	if config == nil {
		err = fmt.Errorf("配置不存在")
		g.Log().Error(ctx, err)
		return
	}

	err = mailerLogicInstance.SendTemplateMail(ctx, config.Value, TemplateData{
		Content: content,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

// convertDeviceToLogic 将实体转换为逻辑模型
func (d *device) convertDeviceToLogic(in *entity.TDevice) (out *model.Device) {
	return &model.Device{
		ID:              in.ID,
		Name:            in.Name,
		DeviceKey:       in.DeviceKey,
		OrgID:           in.OrgID,
		CreatorID:       in.CreatorID,
		Enabled:         model.DeviceStatus(in.Enabled),
		OnlineStatus:    model.DeviceOnlineStatus(in.OnlineStatus),
		Location:        in.Location,
		Description:     in.Description,
		LastOnlineTime:  in.LastOnlineTime,
		LastOfflineTime: in.LastOfflineTime,
		CreatedAt:       in.CreatedAt,
		UpdatedAt:       in.UpdatedAt,
	}
}

func (d *device) DeviceCountHandler(ctx context.Context, params interface{}) (success bool, result string) {
	orgID := params.(string)
	total, err := dao.Device.Ctx(ctx).Where(dao.Device.Columns().OrgID, orgID).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return false, err.Error()
	}

	return true, fmt.Sprintf("设备总数: %d", total)
}
