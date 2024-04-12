package featureimplementation1

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFeature),
)
