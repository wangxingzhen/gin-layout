package logx

import (
	"gin-layout/internal/conf"
	logs "github.com/sirupsen/logrus"
)

// NewLogger 是用于创建logrus实例的函数
func NewLogger(appConf *conf.AppConfig) *logs.Logger {
	logger := logs.New()

	// 设置logrus输出的格式
	logger.Formatter = &logs.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}

	if appConf.Env != conf.EnvLong && appConf.Env != conf.EnvProd {
		logger.SetLevel(logs.DebugLevel)
	}

	return logger
}
