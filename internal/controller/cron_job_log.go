package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
)

type cronJobLogController struct{}

var CronJobLog = cronJobLogController{}

func (c *cronJobLogController) GetJobLogs(ctx context.Context, req *v1.GetJobLogsReq) (res *v1.GetJobLogsRes, err error) {
	out, currentPage, total, err := service.CronJobLog().GetLogs(ctx, req)
	if err != nil {
		return
	}

	res = &v1.GetJobLogsRes{
		PageRes: model.PageRes{
			Total:       total,
			CurrentPage: currentPage,
		},
	}
	for _, v := range out {
		res.List = append(res.List, &v1.CronJobLog{
			ID:        v.ID,
			JobID:     v.JobID,
			JobName:   v.JobName,
			Status:    int(v.ExecuteStatus),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
			Duration:  v.Duration,
			CreatedAt: v.CreatedAt,
		})
	}
	return
}
