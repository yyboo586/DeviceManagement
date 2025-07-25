package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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

/*
func (d *deviceLog) Add(ctx context.Context, in *v1.AddDeviceLogReq) (err error) {
	content, err := json.Marshal(in.Content)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	insertData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:     in.OrgID,
		dao.DeviceLog.Columns().DeviceID:  in.DeviceID,
		dao.DeviceLog.Columns().DeviceKey: in.DeviceKey,
		dao.DeviceLog.Columns().Type:      in.Type,
		dao.DeviceLog.Columns().Content:   string(content),
		dao.DeviceLog.Columns().CreatedAt: in.CreatedAt,
	}
	_, err = dao.DeviceLog.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}
*/

func (d *deviceLog) List(ctx context.Context, req *v1.DeviceLogListReq) (out []*model.DeviceLog, pageRes *model.PageRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = model.DefaultPageSize
	}

	m := dao.DeviceLog.Ctx(ctx).Where(dao.DeviceLog.Columns().OrgID, req.OrgID)

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
	err = m.Page(req.Page, req.PageSize).OrderDesc(dao.DeviceLog.Columns().CreatedAt).Scan(&logEntities)
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
	err = json.Unmarshal([]byte(in.Content), &content)
	if err != nil {
		return
	}

	out = &model.DeviceLog{
		ID:        in.ID,
		OrgID:     in.OrgID,
		DeviceID:  in.DeviceID,
		DeviceKey: in.DeviceKey,
		Type:      model.DeviceLogType(in.Type),
		Content:   &content,
		CreatedAt: gtime.New(time.Unix(in.CreatedAt, 0)),
	}
	return
}
