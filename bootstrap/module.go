package bootstrap

import (
    "deepl-poc/domain"
    "deepl-poc/pkg"
    "deepl-poc/seeds"

    "go.uber.org/fx"
)

var CommonModules = fx.Module("common",
    fx.Options(
        pkg.Module,
        seeds.Module,
        domain.Module,
    ),
)
