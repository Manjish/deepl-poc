package domain

import (
	"deepl-poc/domain/middlewares"

	"deepl-poc/domain/translation"

	"go.uber.org/fx"
)

var Module = fx.Options(
	middlewares.Module,
	translation.Module,
)
