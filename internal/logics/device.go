package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"database/sql"
	"errors"
)

type device struct{}

func NewDevice() *device {
	return &device{}
}

func (d *device) Add(ctx context.Context, in *model.Device) (id int64, err error) {
	insertData := map[string]interface{}{
		dao.Device.Columns().Name:        in.Name,
		dao.Device.Columns().DeviceKey:   in.DeviceKey,
		dao.Device.Columns().OrgID:       in.OrgID,
		dao.Device.Columns().Enabled:     model.DeviceStatusEnabled,
		dao.Device.Columns().Location:    in.Location,
		dao.Device.Columns().Description: in.Description,
	}

	result, err := dao.Device.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		return 0, err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (d *device) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.Device.Ctx(ctx).WhereIn(dao.Device.Columns().ID, ids).Delete()
	return
}

func (d *device) Update(ctx context.Context, in *model.Device) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().Name:        in.Name,
		dao.Device.Columns().Location:    in.Location,
		dao.Device.Columns().Description: in.Description,
	}

	_, err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().ID, in.ID).Data(updateData).Update()
	if err != nil {
		return err
	}

	return
}

func (d *device) Get(ctx context.Context, id int64) (out *model.Device, err error) {
	var device entity.TDevice
	err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().ID, id).Scan(&device)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("设备不存在")
		}
		return nil, err
	}

	out = d.convertDeviceToLogic(&device)
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
		return
	}

	// 获取列表
	var devices []*entity.TDevice
	err = m.Page(page.Page, page.PageSize).Scan(&devices)
	if err != nil {
		return
	}

	for _, device := range devices {
		out = append(out, d.convertDeviceToLogic(device))
	}
	pageRes = &model.PageRes{
		Total:   total,
		Current: page.Page,
	}
	return
}

func (d *device) Enable(ctx context.Context, ids []int64) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().Enabled: model.DeviceStatusEnabled,
	}

	_, err = dao.Device.Ctx(ctx).WhereIn(dao.Device.Columns().ID, ids).Data(updateData).Update()
	return
}

func (d *device) Disable(ctx context.Context, ids []int64) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().Enabled: model.DeviceStatusDisabled,
	}

	_, err = dao.Device.Ctx(ctx).WhereIn(dao.Device.Columns().ID, ids).Data(updateData).Update()
	return
}

func (d *device) Online(ctx context.Context, in *v1.DeviceOnlineReq) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().OnlineStatus: model.DeviceOnlineStatusOnline,
	}

	for _, device := range in.Devices {
		// 更新设备状态
		_, err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().ID, device.DeviceID).Data(updateData).Update()
		if err != nil {
			return err
		}

		// 记录上线日志
		logReq := &model.DeviceOnlineLogReq{
			DeviceID:     device.DeviceID,
			DeviceKey:    device.DeviceKey,
			OrgID:        in.OrgID,
			EventType:    model.DeviceOnlineEventOnline,
			OnlineStatus: model.DeviceOnlineStatusOnline,
			Reason:       "设备上线",
		}

		deviceOnlineLogInstance.Add(ctx, logReq)
	}

	return
}

func (d *device) Offline(ctx context.Context, in *v1.DeviceOfflineReq) (err error) {
	updateData := map[string]interface{}{
		dao.Device.Columns().OnlineStatus: model.DeviceOnlineStatusOffline,
	}

	deviceOnlineLog := NewDeviceOnlineLog()
	for _, device := range in.Devices {
		// 更新设备状态
		_, err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().ID, device.DeviceID).Data(updateData).Update()
		if err != nil {
			return err
		}

		// 记录下线日志
		logReq := &model.DeviceOnlineLogReq{
			DeviceID:     device.DeviceID,
			DeviceKey:    device.DeviceKey,
			OrgID:        in.OrgID,
			EventType:    model.DeviceOnlineEventOffline,
			OnlineStatus: model.DeviceOnlineStatusOffline,
			Reason:       "设备下线",
		}
		deviceOnlineLog.Add(ctx, logReq)
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
