package aws

import (
	"errors"
	"fmt"

	"github.com/yolo-sh/aws-cloud-provider/config"
	"github.com/yolo-sh/aws-cloud-provider/service"
	"github.com/yolo-sh/aws-cloud-provider/userconfig"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/yolo/entities"
)

type AWSViewableErrorBuilder struct {
	presenters.YoloViewableErrorBuilder
}

func NewAWSViewableErrorBuilder() AWSViewableErrorBuilder {
	return AWSViewableErrorBuilder{}
}

func (a AWSViewableErrorBuilder) Build(err error) (viewableError *presenters.ViewableError) {
	viewableError = &presenters.ViewableError{}

	if errors.Is(err, entities.ErrYoloNotInstalled) {
		viewableError.Title = "Yolo not installed"
		viewableError.Message = "Yolo is not installed in this region on this AWS account.\n\n" +
			"Please double check the passed credentials and region."

		return
	}

	if errors.Is(err, entities.ErrUninstallExistingEnvs) {
		viewableError.Title = "Existing environments"
		viewableError.Message = "All environments need to be removed before uninstalling Yolo."

		return
	}

	if errors.Is(err, userconfig.ErrMissingConfig) {
		viewableError.Title = "No AWS account found"
		viewableError.Message = fmt.Sprintf(`An AWS account can be configured:

  - by setting the "%s", "%s" and "%s" environment variables.
		
  - by installing the AWS CLI and running "aws configure".`,
			userconfig.AWSAccessKeyIDEnvVar,
			userconfig.AWSSecretAccessKeyEnvVar,
			userconfig.AWSRegionEnvVar,
		)

		return
	}

	if errors.Is(err, userconfig.ErrMissingAccessKeyInEnv) {
		viewableError.Title = "Missing environment variable"
		viewableError.Message = fmt.Sprintf(
			"The environment variable \"%s\" needs to be set.",
			userconfig.AWSAccessKeyIDEnvVar,
		)

		return
	}

	if errors.Is(err, userconfig.ErrMissingSecretInEnv) {
		viewableError.Title = "Missing environment variable"
		viewableError.Message = fmt.Sprintf(
			"The environment variable \"%s\" needs to be set.",
			userconfig.AWSSecretAccessKeyEnvVar,
		)

		return
	}

	if errors.Is(err, userconfig.ErrMissingRegionInEnv) ||
		errors.Is(err, userconfig.ErrMissingRegionInFiles) {

		viewableError.Title = "Missing region"
		viewableError.Message = fmt.Sprintf(
			"A region needs to be specified by setting the \"%s\" environment variable or by using the \"--region\" flag.",
			userconfig.AWSRegionEnvVar,
		)

		return
	}

	if typedError, ok := err.(userconfig.ErrProfileNotFound); ok {
		viewableError.Title = "Configuration profile not found"
		viewableError.Message = fmt.Sprintf(
			"The profile \"%s\" was not found in your AWS configuration.\n\n(Searched in \"%s\" and \"%s\").",
			typedError.Profile,
			typedError.CredentialsFilePath,
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

	if typedError, ok := err.(config.ErrInvalidAccessKeyID); ok {
		viewableError.Title = "Invalid access key ID"
		viewableError.Message = fmt.Sprintf(
			"The access key ID \"%s\" is invalid.",
			typedError.AccessKeyID,
		)

		return
	}

	if typedError, ok := err.(config.ErrInvalidSecretAccessKey); ok {
		viewableError.Title = "Invalid secret access key"
		viewableError.Message = fmt.Sprintf(
			"The secret access key \"%s\" is invalid.",
			typedError.SecretAccessKey,
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

	if typedError, ok := err.(service.ErrInvalidInstanceTypeArch); ok {
		viewableError.Title = "Unsupported instance type"
		viewableError.Message = fmt.Sprintf(
			"The instance type \"%s\" is not supported by Yolo.\n\n"+
				"Only on-demand linux instances with EBS and \"%s\" architectures are supported.",
			typedError.InstanceType,
			typedError.SupportedArchs,
		)

		return
	}

	viewableError = a.YoloViewableErrorBuilder.Build(err)
	return
}
