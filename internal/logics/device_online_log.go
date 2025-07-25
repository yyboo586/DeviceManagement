package logics

import (
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"sync"

	"github.com/gogf/gf/v2/os/gtime"
)

var (
	deviceOnlineLogOnce     sync.Once
	deviceOnlineLogInstance *deviceOnlineLog
)

type deviceOnlineLog struct{}

func NewDeviceOnlineLog() *deviceOnlineLog {
	deviceOnlineLogOnce.Do(func() {
		deviceOnlineLogInstance = &deviceOnlineLog{}
	})
	return deviceOnlineLogInstance
}

// Add 添加设备上下线日志
func (d *deviceOnlineLog) Add(ctx context.Context, in *model.DeviceOnlineLogReq) (err error) {
	insertData := map[string]interface{}{
		dao.DeviceOnlineLog.Columns().DeviceID:     in.DeviceID,
		dao.DeviceOnlineLog.Columns().DeviceKey:    in.DeviceKey,
		dao.DeviceOnlineLog.Columns().OrgID:        in.OrgID,
		dao.DeviceOnlineLog.Columns().EventType:    in.EventType,
		dao.DeviceOnlineLog.Columns().OnlineStatus: in.OnlineStatus,
		dao.DeviceOnlineLog.Columns().IPAddress:    in.IPAddress,
		dao.DeviceOnlineLog.Columns().ClientID:     in.ClientID,
		dao.DeviceOnlineLog.Columns().Reason:       in.Reason,
		dao.DeviceOnlineLog.Columns().Duration:     in.Duration,
	}

	_, err = dao.DeviceOnlineLog.Ctx(ctx).Data(insertData).Insert()
	return
}

// List 获取设备上下线日志列表
func (d *deviceOnlineLog) List(ctx context.Context, req *model.DeviceOnlineLogListReq) (out []*model.DeviceOnlineLog, pageRes *model.PageRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = model.DefaultPageSize
	}

	m := dao.DeviceOnlineLog.Ctx(ctx).Where(dao.DeviceOnlineLog.Columns().OrgID, req.OrgID)

	// 按设备ID过滤
	if req.DeviceID > 0 {
		m = m.Where(dao.DeviceOnlineLog.Columns().DeviceID, req.DeviceID)
	}

	// 按事件类型过滤
	if req.EventType > 0 {
		m = m.Where(dao.DeviceOnlineLog.Columns().EventType, req.EventType)
	}

	// 按时间范围过滤
	if req.StartTime != "" {
		m = m.WhereGTE(dao.DeviceOnlineLog.Columns().CreatedAt, req.StartTime)
	}
	if req.EndTime != "" {
		m = m.WhereLTE(dao.DeviceOnlineLog.Columns().CreatedAt, req.EndTime)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return
	}

	// 获取列表
	var logs []*entity.TDeviceOnlineLog
	err = m.Page(req.Page, req.PageSize).OrderDesc(dao.DeviceOnlineLog.Columns().CreatedAt).Scan(&logs)
	if err != nil {
		return
	}

	for _, log := range logs {
		out = append(out, d.convertDeviceOnlineLogToLogic(log))
	}

	pageRes = &model.PageRes{
		Total:   total,
		Current: req.Page,
	}
	return
}

// GetDeviceOnlineDuration 获取设备在线时长
func (d *deviceOnlineLog) GetDeviceOnlineDuration(ctx context.Context, deviceID int64, startTime, endTime *gtime.Time) (duration int64, err error) {
	// 获取指定时间范围内的上线记录
	onlineLogs, err := d.getDeviceOnlineLogs(ctx, deviceID, model.DeviceOnlineEventOnline, startTime, endTime)
	if err != nil {
		return 0, err
	}

	// 获取指定时间范围内的下线记录
	offlineLogs, err := d.getDeviceOnlineLogs(ctx, deviceID, model.DeviceOnlineEventOffline, startTime, endTime)
	if err != nil {
		return 0, err
	}

	// 计算在线时长
	duration = d.calculateOnlineDuration(onlineLogs, offlineLogs, startTime, endTime)
	return
}

// getDeviceOnlineLogs 获取设备上下线日志
func (d *deviceOnlineLog) getDeviceOnlineLogs(ctx context.Context, deviceID int64, eventType model.DeviceOnlineEventType, startTime, endTime *gtime.Time) (logs []*entity.TDeviceOnlineLog, err error) {
	m := dao.DeviceOnlineLog.Ctx(ctx).
		Where(dao.DeviceOnlineLog.Columns().DeviceID, deviceID).
		Where(dao.DeviceOnlineLog.Columns().EventType, eventType)

	if startTime != nil {
		m = m.WhereGTE(dao.DeviceOnlineLog.Columns().CreatedAt, startTime)
	}
	if endTime != nil {
		m = m.WhereLTE(dao.DeviceOnlineLog.Columns().CreatedAt, endTime)
	}

	err = m.OrderAsc(dao.DeviceOnlineLog.Columns().CreatedAt).Scan(&logs)
	return
}

// calculateOnlineDuration 计算在线时长
func (d *deviceOnlineLog) calculateOnlineDuration(onlineLogs, offlineLogs []*entity.TDeviceOnlineLog, startTime, endTime *gtime.Time) int64 {
	var totalDuration int64

	// 如果只有上线记录，没有下线记录，说明设备一直在线
	if len(onlineLogs) > 0 && len(offlineLogs) == 0 {
		lastOnlineTime := onlineLogs[len(onlineLogs)-1].CreatedAt
		if endTime != nil {
			totalDuration = endTime.Timestamp() - lastOnlineTime.Timestamp()
		}
		return totalDuration
	}

	// 如果只有下线记录，没有上线记录，说明设备一直离线
	if len(onlineLogs) == 0 && len(offlineLogs) > 0 {
		return 0
	}

	// 计算每次上线的时长
	for _, onlineLog := range onlineLogs {
		var offlineTime *gtime.Time

		// 找到对应的下线时间
		for _, offlineLog := range offlineLogs {
			if offlineLog.CreatedAt.After(onlineLog.CreatedAt) {
				offlineTime = offlineLog.CreatedAt
				break
			}
		}

		// 如果没有找到下线时间，说明设备还在线
		if offlineTime == nil {
			if endTime != nil {
				totalDuration += endTime.Timestamp() - onlineLog.CreatedAt.Timestamp()
			}
		} else {
			totalDuration += offlineTime.Timestamp() - onlineLog.CreatedAt.Timestamp()
		}
	}

	return totalDuration
}

// convertDeviceOnlineLogToLogic 将实体转换为逻辑模型
func (d *deviceOnlineLog) convertDeviceOnlineLogToLogic(in *entity.TDeviceOnlineLog) (out *model.DeviceOnlineLog) {
	return &model.DeviceOnlineLog{
		ID:           in.ID,
		DeviceID:     in.DeviceID,
		DeviceKey:    in.DeviceKey,
		OrgID:        in.OrgID,
		EventType:    model.DeviceOnlineEventType(in.EventType),
		OnlineStatus: model.DeviceOnlineStatus(in.OnlineStatus),
		IPAddress:    in.IPAddress,
		ClientID:     in.ClientID,
		Reason:       in.Reason,
		Duration:     in.Duration,
		CreatedAt:    in.CreatedAt,
	}
}
