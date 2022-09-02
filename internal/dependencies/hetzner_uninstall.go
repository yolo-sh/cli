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

func ProvideHetznerUninstallFeature(yoloConfigDir, region, context string) features.UninstallFeature {
	return provideHetznerUninstallFeature(
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

func provideHetznerUninstallFeature(
	userConfigEnvVarsResolverOpts hetznerProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts hetznerProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts hetznerCLI.UserConfigLocalResolverOpts,
	serviceBuilderOpts hetznerProviderService.BuilderOpts,
) features.UninstallFeature {
	panic(
		wire.Build(
			viewSet,
			hetznerServiceBuilderSet,
			hetznerViewableErrorBuilder,

			stepperSet,

			wire.Bind(new(features.UninstallOutputHandler), new(featuresCLI.UninstallOutputHandler)),
			featuresCLI.NewUninstallOutputHandler,

			wire.Bind(new(featuresCLI.UninstallPresenter), new(presenters.UninstallPresenter)),
			presenters.NewUninstallPresenter,

			wire.Bind(new(presenters.UninstallViewer), new(views.UninstallView)),
			views.NewUninstallView,

			features.NewUninstallFeature,
		),
	)
}
