package service

type ISchedulerService interface {
	Start() error
	Stop()
}

var localScheduler ISchedulerService

func Scheduler() ISchedulerService {
	return localScheduler
}

func RegisterScheduler(i ISchedulerService) {
	localScheduler = i
}
