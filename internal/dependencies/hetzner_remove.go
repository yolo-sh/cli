// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	featuresCLI "github.com/yolo-sh/cli/internal/features"
	hetznerCLI "github.com/yolo-sh/cli/internal/hetzner"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/views"
	hetznerProviderService "github.com/yolo-sh/hetzner-cloud-provider/service"
	hetznerProviderUserConfig "github.com/yolo-sh/hetzner-cloud-provider/userconfig"
	"github.com/yolo-sh/yolo/features"
)

func ProvideHetznerRemoveFeature(yoloConfigDir, region, context string) features.RemoveFeature {
	return provideHetznerRemoveFeature(
		hetznerProviderUserConfig.EnvVarsResolverOpts{
			Region: region,
		},

		hetznerProviderUserConfig.FilesResolverOpts{
			Region:  region,
			Context: context,
		},

		hetznerCLI.UserConfigLocalResolverOpts{
			Context: context,
		},

		hetznerProviderService.BuilderOpts{
			YoloConfigDir: yoloConfigDir,
		},
	)
}

func provideHetznerRemoveFeature(
	userConfigEnvVarsResolverOpts hetznerProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts hetznerProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts hetznerCLI.UserConfigLocalResolverOpts,
	serviceBuilderOpts hetznerProviderService.BuilderOpts,
) features.RemoveFeature {
	panic(
		wire.Build(
			viewSet,
			hetznerServiceBuilderSet,
			hetznerViewableErrorBuilder,

			stepperSet,

			wire.Bind(new(features.RemoveOutputHandler), new(featuresCLI.RemoveOutputHandler)),
			featuresCLI.NewRemoveOutputHandler,

			wire.Bind(new(featuresCLI.RemovePresenter), new(presenters.RemovePresenter)),
			presenters.NewRemovePresenter,

			wire.Bind(new(presenters.RemoveViewer), new(views.RemoveView)),
			views.NewRemoveView,

			features.NewRemoveFeature,
		),
	)
}
