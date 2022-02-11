// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"sweet/internal/autowire"
	"sweet/internal/biz"
	"sweet/internal/conf"
	"sweet/internal/data"
	"sweet/internal/server"
	"sweet/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"sweet/internal/autowire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet,autowire.ProviderSet, newApp))
}
