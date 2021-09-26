// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lynndotconfig/gkgo2/internal/biz"
	"github.com/lynndotconfig/gkgo2/internal/conf"
	"github.com/lynndotconfig/gkgo2/internal/data"
	"github.com/lynndotconfig/gkgo2/internal/server"
	"github.com/lynndotconfig/gkgo2/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
