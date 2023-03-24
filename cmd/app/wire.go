//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gin-layout/internal/biz"
	"gin-layout/internal/conf"
	"gin-layout/internal/data"
	"gin-layout/internal/router"
	"gin-layout/internal/service"
	"gin-layout/pkg/logx"
	"github.com/google/wire"
)

// initApp init app application.
func initApp(appConfig *conf.AppConfig) (*App, func(), error) {
	panic(wire.Build(logx.NewLogger, data.ProviderSet, biz.ProviderSet, service.ProviderSet, router.ProviderSet, newApp))
}
