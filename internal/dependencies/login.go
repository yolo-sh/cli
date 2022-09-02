// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/yolo-sh/cli/internal/features"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/views"
)

func ProvideLoginFeature() features.LoginFeature {
	panic(
		wire.Build(
			viewSet,
			yoloViewableErrorBuilder,

			loggerSet,

			browserManagerSet,

			userConfigManagerSet,

			sleeperSet,

			githubManagerSet,

			wire.Bind(new(features.LoginPresenter), new(presenters.LoginPresenter)),
			presenters.NewLoginPresenter,

			wire.Bind(new(presenters.LoginViewer), new(views.LoginView)),
			views.NewLoginView,

			features.NewLoginFeature,
		),
	)
}
