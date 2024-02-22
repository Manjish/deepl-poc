package pkg

import (
	"deepl-poc/pkg/framework"
	"deepl-poc/pkg/infrastructure"
	"deepl-poc/pkg/middlewares"
	"deepl-poc/pkg/services"

	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	framework.Module,
	services.Module,
	middlewares.Module,
	infrastructure.Module,
)
