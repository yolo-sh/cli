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

func ProvideHetznerOpenPortFeature(yoloConfigDir, region, context string) features.OpenPortFeature {
	return provideHetznerOpenPortFeature(
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

func provideHetznerOpenPortFeature(
	userConfigEnvVarsResolverOpts hetznerProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts hetznerProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts hetznerCLI.UserConfigLocalResolverOpts,
	serviceBuilderOpts hetznerProviderService.BuilderOpts,
) features.OpenPortFeature {
	panic(
		wire.Build(
			viewSet,
			hetznerServiceBuilderSet,
			hetznerViewableErrorBuilder,

			stepperSet,

			wire.Bind(new(features.OpenPortOutputHandler), new(featuresCLI.OpenPortOutputHandler)),
			featuresCLI.NewOpenPortOutputHandler,

			wire.Bind(new(featuresCLI.OpenPortPresenter), new(presenters.OpenPortPresenter)),
			presenters.NewOpenPortPresenter,

			wire.Bind(new(presenters.OpenPortViewer), new(views.OpenPortView)),
			views.NewOpenPortView,

			features.NewOpenPortFeature,
		),
	)
}
