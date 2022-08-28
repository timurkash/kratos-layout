// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/timurkash/kratos-layout/internal/biz"
	"github.com/timurkash/kratos-layout/internal/conf"
	"github.com/timurkash/kratos-layout/internal/outside/data"
	"github.com/timurkash/kratos-layout/internal/server"
	"github.com/timurkash/kratos-layout/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, jwks *conf.Jwks, confData *conf.Data, business *conf.Business, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase, err := biz.NewGreeterUsecase(greeterRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	grpcServer, err := server.NewGRPCServer(confServer, jwks, greeterService, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewHTTPServer(confServer, greeterService, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
