//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/timurkash/kratos-layout/internal/biz"
	"github.com/timurkash/kratos-layout/internal/conf"
	"github.com/timurkash/kratos-layout/internal/outside/data"
	"github.com/timurkash/kratos-layout/internal/server"
	"github.com/timurkash/kratos-layout/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Business, log.Logger) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			service.ProviderSet,
			biz.ProviderSet,
			data.ProviderSet,
			newApp,
		),
	)
}
