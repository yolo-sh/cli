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

func ProvideHetznerClosePortFeature(yoloConfigDir, region, context string) features.ClosePortFeature {
	return provideHetznerClosePortFeature(
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

func provideHetznerClosePortFeature(
	userConfigEnvVarsResolverOpts hetznerProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts hetznerProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts hetznerCLI.UserConfigLocalResolverOpts,
	serviceBuilderOpts hetznerProviderService.BuilderOpts,
) features.ClosePortFeature {
	panic(
		wire.Build(
			viewSet,
			hetznerServiceBuilderSet,
			hetznerViewableErrorBuilder,

			stepperSet,

			wire.Bind(new(features.ClosePortOutputHandler), new(featuresCLI.ClosePortOutputHandler)),
			featuresCLI.NewClosePortOutputHandler,

			wire.Bind(new(featuresCLI.ClosePortPresenter), new(presenters.ClosePortPresenter)),
			presenters.NewClosePortPresenter,

			wire.Bind(new(presenters.ClosePortViewer), new(views.ClosePortView)),
			views.NewClosePortView,

			features.NewClosePortFeature,
		),
	)
}
