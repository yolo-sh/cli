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

func ProvideHetznerEditFeature(yoloConfigDir, region, context string) features.EditFeature {
	return provideHetznerEditFeature(
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

func provideHetznerEditFeature(
	userConfigEnvVarsResolverOpts hetznerProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts hetznerProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts hetznerCLI.UserConfigLocalResolverOpts,
	serviceBuilderOpts hetznerProviderService.BuilderOpts,
) features.EditFeature {
	panic(
		wire.Build(
			viewSet,
			hetznerServiceBuilderSet,
			hetznerViewableErrorBuilder,

			stepperSet,

			vscodeProcessManagerSet,

			vscodeExtensionsManagerSet,

			wire.Bind(new(features.EditOutputHandler), new(featuresCLI.EditOutputHandler)),
			featuresCLI.NewEditOutputHandler,

			wire.Bind(new(featuresCLI.EditPresenter), new(presenters.EditPresenter)),
			presenters.NewEditPresenter,

			wire.Bind(new(presenters.EditViewer), new(views.EditView)),
			views.NewEditView,

			features.NewEditFeature,
		),
	)
}
