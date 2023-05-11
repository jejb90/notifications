//

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"notifications/internal/ctx"
	"notifications/repositories"
)

var stdSet = wire.NewSet(
	repositories.NewGatewayRepository,
	wire.Bind(new(ctx.Gateway), new(*repositories.Gateway)),
	ctx.NewHandler,
)
