package hetzner

import (
	"errors"
	"fmt"

	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/hetzner-cloud-provider/config"
	"github.com/yolo-sh/hetzner-cloud-provider/service"
	"github.com/yolo-sh/hetzner-cloud-provider/userconfig"
	"github.com/yolo-sh/yolo/entities"
)

type HetznerViewableErrorBuilder struct {
	presenters.YoloViewableErrorBuilder
}

func NewHetznerViewableErrorBuilder() HetznerViewableErrorBuilder {
	return HetznerViewableErrorBuilder{}
}

func (h HetznerViewableErrorBuilder) Build(err error) (viewableError *presenters.ViewableError) {
	viewableError = &presenters.ViewableError{}

	if errors.Is(err, entities.ErrYoloNotInstalled) {
		viewableError.Title = "Yolo not installed"
		viewableError.Message = "Yolo is not installed in this region on this Hetzner account.\n\n" +
			"Please double check the passed API token and region."

		return
	}

	if errors.Is(err, entities.ErrUninstallExistingEnvs) {
		viewableError.Title = "Existing environments"
		viewableError.Message = "All environments need to be removed before uninstalling Yolo."

		return
	}

	if errors.Is(err, userconfig.ErrMissingConfig) {
		viewableError.Title = "No Hetzner account found"
		viewableError.Message = fmt.Sprintf(`An Hetzner account can be configured:

  - by setting the "%s" and "%s" environment variables.
		
  - by installing the Hetzner CLI and running "hcloud context create <my_project>".`,
			userconfig.HetznerAPITokenEnvVar,
			userconfig.HetznerRegionEnvVar,
		)

		return
	}

	if errors.Is(err, userconfig.ErrMissingRegionInEnv) {
		viewableError.Title = "Missing region"
		viewableError.Message = fmt.Sprintf(
			"A region needs to be specified by setting the \"%s\" environment variable or by using the \"--region\" flag.",
			userconfig.HetznerRegionEnvVar,
		)

		return
	}

	if errors.Is(err, userconfig.ErrMissingRegion) {
		viewableError.Title = "Missing region"
		viewableError.Message = "A region needs to be specified by using the \"--region\" flag."

		return
	}

	if typedError, ok := err.(config.ErrContextNotFound); ok {
		viewableError.Title = "Configuration context not found"
		viewableError.Message = fmt.Sprintf(
			"The context \"%s\" was not found in your Hetzner configuration.\n\n(Searched in \"%s\").",
			typedError.Context,
			typedError.ConfigFilePath,
		)

		return
	}

	if typedError, ok := err.(config.ErrInvalidRegion); ok {
		viewableError.Title = "Invalid region"
		viewableError.Message = fmt.Sprintf(
			"The region \"%s\" is invalid.",
			typedError.Region,
		)

		return
	}

	if typedError, ok := err.(config.ErrInvalidAPIToken); ok {
		viewableError.Title = "Invalid API token"
		viewableError.Message = fmt.Sprintf(
			"The API token \"%s\" is invalid.",
			typedError.APIToken,
		)

		return
	}

	if typedError, ok := err.(service.ErrInvalidInstanceType); ok {
		viewableError.Title = "Invalid instance type"
		viewableError.Message = fmt.Sprintf(
			"The instance type \"%s\" is invalid in the region \"%s\".",
			typedError.InstanceType,
			typedError.Region,
		)

		return
	}

	viewableError = h.YoloViewableErrorBuilder.Build(err)
	return
}
