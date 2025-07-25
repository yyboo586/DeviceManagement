package main

import (
	"DeviceManagement/internal/common"
	"DeviceManagement/internal/controller"
	"DeviceManagement/internal/logics"
	"DeviceManagement/internal/service"
	"net/http"

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

	_ = logics.NewMailer()
	logicsProductCategory := logics.NewProductCategory()
	logicsProduct := logics.NewProduct()
	logicsDevice := logics.NewDevice()
	logicsDeviceLog := logics.NewDeviceLog()
	logicsDeviceConfig := logics.NewDeviceConfig()
	logicsCronJob := logics.NewCronJob()
	logicsCronJobLog := logics.NewCronJobLog()
	logicsScheduler := logics.NewScheduler()
	logicsThirdService := logics.NewThirdService()

	service.RegisterProductCategory(logicsProductCategory)
	service.RegisterProduct(logicsProduct)
	service.RegisterDevice(logicsDevice)
	service.RegisterDeviceLog(logicsDeviceLog)
	service.RegisterDeviceConfig(logicsDeviceConfig)
	service.RegisterCronJob(logicsCronJob)
	service.RegisterCronJobLog(logicsCronJobLog)
	service.RegisterScheduler(logicsScheduler)
	service.RegisterThirdService(logicsThirdService)
	service.RegisterMQService()

	//if err := service.Scheduler().Start(context.Background()); err != nil {
	//	g.Log().Fatalf(context.Background(), "failed to start scheduler: %v", err)
	//}

	s.Group("/api/v1/device-management", func(group *ghttp.RouterGroup) {
		group.Middleware(CORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Middleware(Auth)
		group.Bind(
			controller.ProductCategoryController,
			controller.ProductController,
			controller.Device,
			controller.DeviceLog,
			controller.DeviceConfig,
			controller.CronJob,
			controller.CronJobLog,
		)
	})

	s.Run()
}

func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func Auth(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.Status = http.StatusOK
		r.Response.WriteJson(g.Map{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		r.Exit()
	}

	out, err := service.ThirdService().Introspect(r.Context(), token)
	if err != nil {
		r.Response.Status = http.StatusOK
		r.Response.WriteJson(g.Map{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		r.Exit()
	}

	r.SetCtxVar(common.TokenInspectRes, out)
	r.Middleware.Next()
}
