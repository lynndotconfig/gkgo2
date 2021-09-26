// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lynndotconfig/gkgo2/internal/biz"
	"github.com/lynndotconfig/gkgo2/internal/conf"
	"github.com/lynndotconfig/gkgo2/internal/data"
	"github.com/lynndotconfig/gkgo2/internal/server"
	"github.com/lynndotconfig/gkgo2/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	accountRepo := data.NewAccountRepo(dataData, logger)
	accountUsecase := biz.NewAccountUsecase(accountRepo, logger)
	accountService := service.NewAccountService(accountUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, accountService, logger)
	grpcServer := server.NewGRPCServer(confServer, accountService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
