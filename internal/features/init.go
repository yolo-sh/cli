package features

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/yolo-sh/agent/constants"
	"github.com/yolo-sh/agent/proto"
	"github.com/yolo-sh/cli/internal/agent"
	"github.com/yolo-sh/cli/internal/config"
	cliConstants "github.com/yolo-sh/cli/internal/constants"
	cliEntities "github.com/yolo-sh/cli/internal/entities"
	"github.com/yolo-sh/cli/internal/interfaces"
	"github.com/yolo-sh/yolo/actions"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/features"
)

type InitResponse struct {
	Error   error
	Content InitResponseContent
}

type InitResponseContent struct {
	EnvName            string
	EnvPublicIPAddress string
	EnvAlreadyCreated  bool
}

type InitPresenter interface {
	PresentToView(InitResponse)
}

type InitOutputHandler struct {
	userConfig         interfaces.UserConfigManager
	presenter          InitPresenter
	agentClientBuilder agent.ClientBuilder
	github             interfaces.GitHubManager
	logger             interfaces.Logger
	sshConfig          interfaces.SSHConfigManager
	sshKeys            interfaces.SSHKeysManager
	sshKnownHosts      interfaces.SSHKnownHostsManager
}

func NewInitOutputHandler(
	userConfig interfaces.UserConfigManager,
	presenter InitPresenter,
	agentClientBuilder agent.ClientBuilder,
	github interfaces.GitHubManager,
	logger interfaces.Logger,
	sshConfig interfaces.SSHConfigManager,
	sshKeys interfaces.SSHKeysManager,
	sshKnownHosts interfaces.SSHKnownHostsManager,
) InitOutputHandler {

	return InitOutputHandler{
		userConfig:         userConfig,
		presenter:          presenter,
		agentClientBuilder: agentClientBuilder,
		github:             github,
		logger:             logger,
		sshConfig:          sshConfig,
		sshKnownHosts:      sshKnownHosts,
		sshKeys:            sshKeys,
	}
}

func (i InitOutputHandler) HandleOutput(output features.InitOutput) error {

	stepper := output.Stepper

	handleError := func(err error) error {
		stepper.StopCurrentStep()

		i.presenter.PresentToView(InitResponse{
			Error: err,
		})

		return err
	}

	if output.Error != nil {
		return handleError(output.Error)
	}

	env := output.Content.Env

	envCreated := output.Content.EnvCreated
	envAlreadyCreated := !envCreated

	var envAdditionalProperties *cliEntities.EnvAdditionalProperties

	if len(env.AdditionalPropertiesJSON) > 0 {
		err := json.Unmarshal(
			[]byte(env.AdditionalPropertiesJSON),
			&envAdditionalProperties,
		)

		if err != nil {
			return handleError(err)
		}
	}

	if envAdditionalProperties == nil {
		envAdditionalProperties = &cliEntities.EnvAdditionalProperties{}
	}

	if envCreated {
		stepper.StartTemporaryStep(
			"Building the environment",
		)

		agentClient := i.agentClientBuilder.Build(
			agent.NewDefaultClientConfig(
				[]byte(env.SSHKeyPairPEMContent),
				env.InstancePublicIPAddress,
			),
		)

		err := agentClient.InitInstance(
			&proto.InitInstanceRequest{},
			func(stream agent.InitInstanceStream) error {

				for {
					_, err := stream.Recv()

					if err == io.EOF {
						break
					}

					if err != nil {
						return err
					}
				}

				return nil
			},
		)

		if err != nil {
			return handleError(err)
		}

		resolvedRepository := env.ResolvedRepository

		err = agentClient.BuildAndStartEnv(
			&proto.BuildAndStartEnvRequest{
				EnvNameSlug:          env.GetNameSlug(),
				EnvRepoOwner:         resolvedRepository.Owner,
				EnvRepoName:          resolvedRepository.Name,
				EnvRepoLanguagesUsed: resolvedRepository.LanguagesUsed,
			},
			func(stream agent.BuildAndStartEnvStream) error {
				for {
					reply, err := stream.Recv()

					if err == io.EOF {
						break
					}

					if err != nil {
						return err
					}

					stepper.StopCurrentStep()

					if len(reply.LogLineHeader) > 0 {
						bold := cliConstants.Bold
						blue := cliConstants.Blue
						i.logger.Log(bold(blue("> " + reply.LogLineHeader + "\n")))
					}

					if len(reply.LogLine) > 0 {
						i.logger.LogNoNewline(reply.LogLine)
					}
				}

				return nil
			},
		)

		if err != nil {
			return handleError(err)
		}

		stepper.StartTemporaryStep(
			"Initializing the environment",
		)

		err = agentClient.InitEnv(&proto.InitEnvRequest{
			EnvRepoOwner:         resolvedRepository.Owner,
			EnvRepoName:          resolvedRepository.Name,
			EnvRepoLanguagesUsed: resolvedRepository.LanguagesUsed,
			GithubUserEmail:      i.userConfig.GetString(config.UserConfigKeyGitHubEmail),
			UserFullName:         i.userConfig.GetString(config.UserConfigKeyGitHubFullName),
		}, func(stream agent.InitEnvStream) error {

			for {
				reply, err := stream.Recv()

				if err == io.EOF {
					break
				}

				if err != nil {
					return err
				}

				if reply.GithubSshPublicKeyContent != nil &&
					envAdditionalProperties.GitHubCreatedSSHKeyId == nil {

					sshKeyInGitHub, err := i.github.CreateSSHKey(
						i.userConfig.GetString(config.UserConfigKeyGitHubAccessToken),
						fmt.Sprintf("yolo-%s", env.GetNameSlug()),
						reply.GetGithubSshPublicKeyContent(),
					)

					if err != nil {
						return err
					}

					envAdditionalProperties.GitHubCreatedSSHKeyId = sshKeyInGitHub.ID
					err = env.SetAdditionalPropertiesJSON(envAdditionalProperties)

					if err != nil {
						return err
					}

					err = actions.UpdateEnvInConfig(
						stepper,
						output.Content.CloudService,
						output.Content.YoloConfig,
						output.Content.Cluster,
						env,
					)

					if err != nil {
						return err
					}
				}

				if reply.GithubGpgPublicKeyContent != nil &&
					envAdditionalProperties.GitHubCreatedGPGKeyId == nil {

					gpgKeyInGitHub, err := i.github.CreateGPGKey(
						i.userConfig.GetString(config.UserConfigKeyGitHubAccessToken),
						reply.GetGithubGpgPublicKeyContent(),
					)

					if err != nil {
						return err
					}

					envAdditionalProperties.GitHubCreatedGPGKeyId = gpgKeyInGitHub.ID
					err = env.SetAdditionalPropertiesJSON(envAdditionalProperties)

					if err != nil {
						return err
					}

					err = actions.UpdateEnvInConfig(
						stepper,
						output.Content.CloudService,
						output.Content.YoloConfig,
						output.Content.Cluster,
						env,
					)

					if err != nil {
						return err
					}
				}
			}

			return nil
		})

		if err != nil {
			return handleError(err)
		}
	}

	if !envAlreadyCreated {
		err := output.Content.SetEnvAsCreated()

		if err != nil {
			return handleError(err)
		}
	}

	stepper.StartTemporaryStepWithoutNewLine(
		"Updating your local SSH configuration",
	)

	sshPEMPath, err := i.sshKeys.CreateOrReplacePEM(
		env.GetSSHKeyPairName(),
		env.SSHKeyPairPEMContent,
	)

	if err != nil {
		return handleError(err)
	}

	sshServerListenPort, err := strconv.ParseInt(
		constants.SSHServerListenPort,
		10,
		64,
	)

	if err != nil {
		return handleError(err)
	}

	sshConfigHostKey := env.Name

	err = i.sshConfig.AddOrReplaceHost(
		sshConfigHostKey,
		env.InstancePublicIPAddress,
		sshPEMPath,
		entities.EnvRootUser,
		sshServerListenPort,
	)

	if err != nil {
		return handleError(err)
	}

	for _, sshHostKey := range env.SSHHostKeys {
		err := i.sshKnownHosts.AddOrReplace(
			env.InstancePublicIPAddress,
			sshHostKey.Algorithm,
			sshHostKey.Fingerprint,
		)

		if err != nil {
			return handleError(err)
		}
	}

	stepper.StopCurrentStep()

	i.presenter.PresentToView(InitResponse{
		Content: InitResponseContent{
			EnvName:            env.Name,
			EnvPublicIPAddress: env.InstancePublicIPAddress,
			EnvAlreadyCreated:  envAlreadyCreated,
		},
	})

	return nil
}
