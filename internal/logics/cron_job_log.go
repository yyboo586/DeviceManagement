package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
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
	if in.PageReq.Page <= 0 {
		in.PageReq.Page = 1
	}
	currentPage = in.PageReq.Page
	if in.PageReq.PageSize <= 0 {
		in.PageReq.PageSize = 10
	}

	m := dao.CronJobLog.Ctx(ctx).Where(dao.CronJobLog.Columns().OrgID, in.OrgID)

	if in.Name != "" {
		m = m.Where(dao.CronJobLog.Columns().JobID, in.Name)
	}

	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	entities := make([]*entity.TCronJobLog, 0)
	err = m.Page(in.PageReq.Page, in.PageReq.PageSize).Scan(&entities)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, entity := range entities {
		log := &model.CronJobLog{
			ID:            entity.ID,
			JobID:         entity.JobID,
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
