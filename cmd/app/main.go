package main

import (
	"flag"
	"gin-layout/internal/conf"
	"gin-layout/pkg"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

var (
	// APIConfig api的配置
	config  conf.AppConfig
	Version string
)

type App struct {
	conf   *conf.AppConfig
	gin    *gin.Engine
	logger *logs.Logger
}

func init() {
	err := pkg.LoadConfigFor(&config, "app.yml")
	if err != nil {
		logs.Panicf("LoadAppConfig error: %v", err)
	}
}

func newApp(conf *conf.AppConfig, engine *gin.Engine, logs *logs.Logger) *App {
	return &App{conf: conf, gin: engine, logger: logs}
}

func main() {
	var (
		port = flag.String("p", ":8082", "--p=:port")
	)
	flag.Parse()

	app, f, err := initApp(&config)
	if err != nil {
		panic(err)
		return
	}
	defer f()

	if len(Version) > 0 {
		app.logger.Infof("git commit: %v", Version)
	}

	panic(app.Run(*port))
}

// Run start service
func (app *App) Run(port string) error {
	return app.gin.Run(port)
}
