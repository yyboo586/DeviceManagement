package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	cronJobLogOnce     sync.Once
	cronJobLogInstance *cronJobLog
)

type cronJobLog struct {
}

func NewCronJobLog() *cronJobLog {
	cronJobLogOnce.Do(func() {
		cronJobLogInstance = &cronJobLog{}
	})
	return cronJobLogInstance
}

func (c *cronJobLog) GetLogs(ctx context.Context, in *v1.GetJobLogsReq) (out []*model.CronJobLog, currentPage int, total int, err error) {
	operatorInfo := ctx.Value(common.TokenInspectRes).(*model.TokenData)
	if in.PageReq.Page <= 0 {
		in.PageReq.Page = 1
	}
	currentPage = in.PageReq.Page
	if in.PageReq.PageSize <= 0 {
		in.PageReq.PageSize = 10
	}

	m := dao.CronJobLog.Ctx(ctx).Where(dao.CronJobLog.Columns().OrgID, operatorInfo.OrgID)
	if in.JobID != "" {
		m = m.Where(dao.CronJobLog.Columns().JobID, in.JobID)
	}
	if in.ExecuteStatus != "" {
		status := model.GetCronJobExecuteStatusStr(in.ExecuteStatus)
		if status == model.CronJobExecuteStatusUnknown {
			err = fmt.Errorf("请求参数错误: 执行状态不正确")
			return
		}
		m = m.Where(dao.CronJobLog.Columns().ExecuteStatus, status)
	}

	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	entities := make([]*entity.TCronJobLog, 0)
	err = m.Page(in.PageReq.Page, in.PageReq.PageSize).Order(dao.CronJobLog.Columns().CreatedAt, "DESC").Scan(&entities)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, entity := range entities {
		log := &model.CronJobLog{
			ID:            entity.ID,
			JobID:         entity.JobID,
			JobName:       entity.JobName,
			ExecuteStatus: model.CronJobExecuteStatus(entity.ExecuteStatus),
			Result:        entity.Result,
			StartTime:     entity.StartTime,
			EndTime:       entity.EndTime,
			Duration:      entity.Duration,
			CreatedAt:     entity.CreatedAt,
		}
		out = append(out, log)
	}
	return
}
