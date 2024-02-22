package translation

import (
    "go.uber.org/fx"
)

var Module = fx.Module("translation",
	fx.Options(
		fx.Provide(
			NewRepository,
			NewService,
			NewController,
			NewRoute,
		),
		fx.Invoke(RegisterRoute),
	),
)
