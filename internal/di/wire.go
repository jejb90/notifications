//

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"notifications/internal/ctx"
)

func Initialize() (*ctx.Handler, error) {
	wire.Build(stdSet)

	return &ctx.Handler{}, nil
}
