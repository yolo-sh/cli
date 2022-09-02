// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/yolo-sh/cli/internal/entities"
)

func ProvideEnvRepositoryResolver() entities.EnvRepositoryResolver {
	panic(
		wire.Build(
			loggerSet,

			userConfigManagerSet,

			githubManagerSet,

			entities.NewEnvRepositoryResolver,
		),
	)
}
