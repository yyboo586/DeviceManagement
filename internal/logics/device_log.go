package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"encoding/json"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	deviceLogOnce     sync.Once
	deviceLogInstance *deviceLog
)

type deviceLog struct{}

func NewDeviceLog() *deviceLog {
	deviceLogOnce.Do(func() {
		deviceLogInstance = &deviceLog{}
	})
	return deviceLogInstance
}

func (d *deviceLog) List(ctx context.Context, req *v1.DeviceLogListReq) (out []*model.DeviceLog, pageRes *model.PageRes, err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = model.DefaultPageSize
	}

	m := dao.DeviceLog.Ctx(ctx).Where(dao.DeviceLog.Columns().OrgID, operatorInfo.OrgID)
	// 按设备ID过滤
	if req.DeviceID > 0 {
		m = m.Where(dao.DeviceLog.Columns().DeviceID, req.DeviceID)
	}
	// 按事件类型过滤
	if req.Type > 0 {
		m = m.Where(dao.DeviceLog.Columns().Type, req.Type)
	}
	// 按时间范围过滤
	if req.StartTime != nil {
		m = m.WhereGTE(dao.DeviceLog.Columns().CreatedAt, req.StartTime)
	}
	if req.EndTime != nil {
		m = m.WhereLTE(dao.DeviceLog.Columns().CreatedAt, req.EndTime)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 获取列表
	var logEntities []*entity.TDeviceLog
	err = m.Page(req.Page, req.PageSize).Order(dao.DeviceLog.Columns().CreatedAt, "DESC").Scan(&logEntities)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range logEntities {
		var logModel *model.DeviceLog
		logModel, err = d.convertEntityToLogic(v)
		if err != nil {
			g.Log().Error(ctx, err)
			break
		}
		out = append(out, logModel)
	}

	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: req.Page,
	}
	return
}

// convertEntityToLogic 将实体转换为逻辑模型
func (d *deviceLog) convertEntityToLogic(in *entity.TDeviceLog) (out *model.DeviceLog, err error) {
	var content model.DeviceLogContent
	if in.Content != "" {
		err = json.Unmarshal([]byte(in.Content), &content)
		if err != nil {
			return
		}
	}

	out = &model.DeviceLog{
		ID:         in.ID,
		OrgID:      in.OrgID,
		DeviceID:   in.DeviceID,
		DeviceName: in.DeviceName,
		DeviceKey:  in.DeviceKey,
		Type:       model.DeviceLogType(in.Type),
		Content:    &content,
		Timestamp:  in.Timestamp,
		CreatedAt:  in.CreatedAt,
	}
	return
}
