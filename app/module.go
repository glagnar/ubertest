package app

import (
	featureimplementation1 "github.com/glagnar/ubertest/feature/implementation1"
	//featureimplementation2 "github.com/glagnar/ubertest/feature/implementation2"
	"github.com/glagnar/ubertest/program"
	"go.uber.org/fx"
)

var BaseModule = fx.Options(
	featureimplementation1.Module,
	//featureimplementation2.Module,

	fx.Provide(program.NewCoreApplication),
)
