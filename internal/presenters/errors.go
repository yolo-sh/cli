package presenters

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yolo-sh/cli/internal/constants"
	"github.com/yolo-sh/cli/internal/exceptions"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ViewableError struct {
	Title   string
	Message string
}

type ViewableErrorBuilder interface {
	Build(error) *ViewableError
}

type YoloViewableErrorBuilder struct{}

func NewYoloViewableErrorBuilder() YoloViewableErrorBuilder {
	return YoloViewableErrorBuilder{}
}

func (YoloViewableErrorBuilder) Build(err error) (viewableError *ViewableError) {
	viewableError = &ViewableError{}

	if typedError, ok := err.(entities.ErrClusterNotExists); ok {
		viewableError.Title = "Cluster not found"
		viewableError.Message = fmt.Sprintf(
			"The cluster \"%s\" was not found.",
			typedError.ClusterName,
		)

		return
	}

	if typedError, ok := err.(entities.ErrClusterAlreadyExists); ok {
		viewableError.Title = "Cluster already exists"
		viewableError.Message = fmt.Sprintf(
			"The cluster \"%s\" already exists.",
			typedError.ClusterName,
		)

		return
	}

	if typedError, ok := err.(entities.ErrEnvNotExists); ok {
		viewableError.Title = "Environment not found"

		if typedError.ClusterName != entities.DefaultClusterName {
			viewableError.Message = fmt.Sprintf(
				"The environment for \"%s\" was not found in the cluster \"%s\".",
				typedError.EnvName,
				typedError.ClusterName,
			)
			return
		}

		viewableError.Message = fmt.Sprintf(
			"The environment for \"%s\" was not found.",
			typedError.EnvName,
		)
		return
	}

	if errors.Is(err, exceptions.ErrUserNotLoggedIn) {
		viewableError.Title = "GitHub account not connected"
		viewableError.Message = fmt.Sprintf(
			"You must first connect your GitHub account using the command \"yolo login\".\n\n"+
				"Yolo requires the following permissions:\n\n"+
				"  - \"Public SSH keys\" and \"Repositories\" to let you access your repositories from your environments\n\n"+
				"  - \"GPG Keys\" and \"Personal user data\" to configure Git and sign your commits (verified badge)\n\n"+
				"All your data (including the OAuth access token) will only be stored locally (in \"%s\").",
			system.UserConfigFilePath(),
		)

		return
	}

	if typedError, ok := err.(entities.ErrEnvRepositoryNotFound); ok {
		viewableError.Title = "Repository not found"
		viewableError.Message = fmt.Sprintf(
			"The repository \"%s/%s\" was not found.\n\n"+
				"Please double check that this repository exists and that you can access it.",
			typedError.RepoOwner,
			typedError.RepoName,
		)

		return
	}

	if typedError, ok := err.(entities.ErrInitRemovingEnv); ok {
		viewableError.Title = "Invalid environment state"
		viewableError.Message = fmt.Sprintf(
			"The environment for \"%s\" cannot be initialized because it's currently removing.\n\n"+
				"You must wait for the removing process to terminate.",
			typedError.EnvName,
		)

		return
	}

	if typedError, ok := err.(entities.ErrEditRemovingEnv); ok {
		viewableError.Title = "Invalid environment state"
		viewableError.Message = fmt.Sprintf(
			"The environment for \"%s\" cannot be edited because it's currently removing.\n\n"+
				"You must wait for the removing process to terminate.",
			typedError.EnvName,
		)

		return
	}

	if typedError, ok := err.(entities.ErrEditCreatingEnv); ok {
		viewableError.Title = "Invalid environment state"
		viewableError.Message = fmt.Sprintf(
			"The environment for \"%s\" cannot be edited because it's currently creating.\n\n"+
				"You must wait for the creation process to terminate.",
			typedError.EnvName,
		)

		return
	}

	if typedError, ok := err.(entities.ErrInvalidPort); ok {
		viewableError.Title = "Invalid port"
		viewableError.Message = fmt.Sprintf(
			"The port \"%s\" is invalid. Must be a number in the range 1-65535.",
			typedError.InvalidPort,
		)

		return
	}

	if typedError, ok := err.(entities.ErrReservedPort); ok {
		viewableError.Title = "Reserved port"
		viewableError.Message = fmt.Sprintf(
			"The port \"%s\" is reserved for Yolo usage.",
			typedError.ReservedPort,
		)

		return
	}

	if typedError, ok := err.(exceptions.ErrLoginError); ok {
		viewableError.Title = "GitHub connection error"
		viewableError.Message = fmt.Sprintf(
			"An error has occured during the authorization of the Yolo application.\n\n%s",
			typedError.Reason,
		)

		return
	}

	if typedError, ok := err.(exceptions.ErrMissingRequirements); ok {
		viewableError.Title = "Missing requirements"
		viewableError.Message = fmt.Sprintf(
			"The following requirements are missing:\n\n  - %s",
			strings.Join(typedError.MissingRequirements, "\n\n  - "),
		)

		return
	}

	bold := constants.Bold

	if status, ok := status.FromError(err); ok {
		viewableError.Title = "Yolo agent error"

		errorMessage := status.Message()

		if len(errorMessage) >= 2 {
			errorMessage = strings.ToTitle(errorMessage[0:1]) + errorMessage[1:] + "."
		}

		viewableError.Message = errorMessage

		if status.Code() != codes.Unknown {
			viewableError.Message += "\n\n" +
				bold("Error code: ") +
				status.Code().String()
		}

		return
	}

	viewableError.Title = "Unknown error"
	viewableError.Message = fmt.Sprintf(
		"An unknown error occurred.\n\n"+
			"You could try to fix it (using the details below) or open a new issue at: https://github.com/yolo-sh/cli/issues/new\n\n"+
			bold("%s"),
		err.Error(),
	)

	return
}
