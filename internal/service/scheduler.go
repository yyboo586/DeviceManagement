package service

import "context"

type ISchedulerService interface {
	Start(ctx context.Context) error
	Stop()
}

var localScheduler ISchedulerService

func Scheduler() ISchedulerService {
	return localScheduler
}

func RegisterScheduler(i ISchedulerService) {
	localScheduler = i
}
