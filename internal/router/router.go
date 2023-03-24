package router

import (
	"fmt"
	"gin-layout/internal/conf"
	"gin-layout/internal/service"
	"gin-layout/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logs "github.com/sirupsen/logrus"
	"time"
)

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(
	NewRouter,
	NewBeforeHandel,
)

func NewRouter(user *service.UserService, appConfig *conf.AppConfig,
	beforeHandel *RequestBeforeHandel,
	logger *logs.Logger,
) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	if appConfig.Env == conf.EnvProd || appConfig.Env == conf.EnvLong {
		gin.SetMode(gin.ReleaseMode)
	}

	// 日志放入logrus中
	gin.DefaultWriter = logs.NewEntry(logger).WriterLevel(logs.InfoLevel)
	gin.DefaultErrorWriter = logs.NewEntry(logger).WriterLevel(logs.ErrorLevel)

	router := gin.New()

	// 更改gin的log包
	router.Use(GenLogger(logger))
	router.Use(GenGinRecover(), GenGinLogger())

	// example ... start

	test := router.Group("/test")
	{
		test.GET("", ginx.API(user.Test))
		test.POST("/add", ginx.API(user.AddTest))
		test.POST("/tran", ginx.API(user.TranTest, beforeHandel.SuperAdmin))
		test.GET("/list", ginx.API(user.ListTest))
	}
	router.Use(VerifyLogin(user))

	// example ... end

	return router
}

// example ... start

type RequestBeforeHandel struct {
	user *service.UserService
}

func NewBeforeHandel(adUser *service.UserService) *RequestBeforeHandel {
	return &RequestBeforeHandel{
		user: adUser,
	}
}

// SuperAdmin 判断用户是不是管理员
func (r *RequestBeforeHandel) SuperAdmin(req ginx.RequestHandler) ginx.RequestHandler {
	return func(rc *ginx.RequestContext) {
		t := time.Now()
		response, err := r.user.IsAdmin(rc)
		if err != nil {
			ginx.ReturnJSON(rc, response, err)
			rc.GetLogger().WithFields(logs.Fields{
				"serviceName": "RequestBeforeHandel",
				"funcName":    "SuperAdmin",
				"elapsed":     fmt.Sprintf("%.3fms", float64(time.Since(t)/time.Millisecond)),
			}).Errorf("error:%+v", err)
			return
		}
		req(rc)
	}
}

// example ... end
