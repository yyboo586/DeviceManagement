package main

import (
	"DeviceManagement/internal/controller"
	"DeviceManagement/internal/logics"
	"DeviceManagement/internal/service"
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	s := g.Server()
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")

	logicsProductCategory := logics.NewProductCategory()
	logicsProduct := logics.NewProduct()
	logicsThingModelTemplate := logics.NewThingModelTemplate()
	logicsThingModel := logics.NewThingModel()
	logicsDevice := logics.NewDevice()
	logicsDeviceOnlineLog := logics.NewDeviceOnlineLog()

	logicsCronJob := logics.NewCronJob()
	logicsScheduler := logics.NewScheduler()
	logicsThirdService := logics.NewThirdService()

	service.RegisterProductCategory(logicsProductCategory)
	service.RegisterProduct(logicsProduct)
	service.RegisterThingModelTemplate(logicsThingModelTemplate)
	service.RegisterThingModel(logicsThingModel)
	service.RegisterDevice(logicsDevice)
	service.RegisterDeviceLog(logicsDeviceOnlineLog)
	service.RegisterCronJob(logicsCronJob)
	service.RegisterScheduler(logicsScheduler)
	service.RegisterThirdService(logicsThirdService)

	if err := service.Scheduler().Start(context.Background()); err != nil {
		g.Log().Fatalf(context.Background(), "failed to start scheduler: %v", err)
	}

	s.Group("/api/v1/device-management", func(group *ghttp.RouterGroup) {
		group.Middleware(CORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			//controller.ProductCategoryController,
			//controller.ProductController,
			//controller.ThingModelController,
			controller.DeviceController,
			controller.CronJobController,
		)
	})

	s.Run()
}

func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
