// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	hetznerCLI "github.com/yolo-sh/cli/internal/hetzner"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/system"
	hetznerProviderConfig "github.com/yolo-sh/hetzner-cloud-provider/config"
	hetznerProviderService "github.com/yolo-sh/hetzner-cloud-provider/service"
	hetznerProviderUserConfig "github.com/yolo-sh/hetzner-cloud-provider/userconfig"
	"github.com/yolo-sh/yolo/entities"
)

var hetznerViewableErrorBuilder = wire.NewSet(
	wire.Bind(new(presenters.ViewableErrorBuilder), new(hetznerCLI.HetznerViewableErrorBuilder)),
	hetznerCLI.NewHetznerViewableErrorBuilder,
)

var hetznerServiceBuilderSet = wire.NewSet(
	wire.Bind(new(hetznerProviderUserConfig.ContextLoader), new(hetznerProviderConfig.ContextLoader)),
	hetznerProviderConfig.NewContextLoader,

	wire.Bind(new(hetznerCLI.UserConfigFilesResolver), new(hetznerProviderUserConfig.FilesResolver)),
	hetznerProviderUserConfig.NewFilesResolver,

	wire.Bind(new(hetznerProviderUserConfig.EnvVarsGetter), new(system.EnvVars)),
	system.NewEnvVars,

	wire.Bind(new(hetznerCLI.UserConfigEnvVarsResolver), new(hetznerProviderUserConfig.EnvVarsResolver)),
	hetznerProviderUserConfig.NewEnvVarsResolver,

	wire.Bind(new(hetznerProviderService.UserConfigResolver), new(hetznerCLI.UserConfigLocalResolver)),
	hetznerCLI.NewUserConfigLocalResolver,

	wire.Bind(new(hetznerProviderService.UserConfigValidator), new(hetznerProviderConfig.UserConfigValidator)),
	hetznerProviderConfig.NewUserConfigValidator,

	wire.Bind(new(entities.CloudServiceBuilder), new(hetznerProviderService.Builder)),
	hetznerProviderService.NewBuilder,
)
