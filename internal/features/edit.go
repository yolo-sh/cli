package features

import (
	"github.com/yolo-sh/agent/constants"
	"github.com/yolo-sh/cli/internal/interfaces"
	"github.com/yolo-sh/yolo/features"
)

type EditResponse struct {
	Error   error
	Content EditResponseContent
}

type EditResponseContent struct {
	EnvName string
}

type EditPresenter interface {
	PresentToView(EditResponse)
}

type EditOutputHandler struct {
	presenter        EditPresenter
	vscodeProcess    interfaces.VSCodeProcessManager
	vscodeExtensions interfaces.VSCodeExtensionsManager
}

func NewEditOutputHandler(
	presenter EditPresenter,
	vscodeProcess interfaces.VSCodeProcessManager,
	vscodeExtensions interfaces.VSCodeExtensionsManager,
) EditOutputHandler {

	return EditOutputHandler{
		presenter:        presenter,
		vscodeProcess:    vscodeProcess,
		vscodeExtensions: vscodeExtensions,
	}
}

func (e EditOutputHandler) HandleOutput(output features.EditOutput) error {
	stepper := output.Stepper

	handleError := func(err error) error {
		stepper.StopCurrentStep()

		e.presenter.PresentToView(EditResponse{
			Error: err,
		})

		return err
	}

	if output.Error != nil {
		return handleError(output.Error)
	}

	stepper.StartTemporaryStepWithoutNewLine(
		"Installing Visual Studio Code Remote - SSH extension",
	)

	_, err := e.vscodeExtensions.Install("ms-vscode-remote.remote-ssh")

	if err != nil {
		return handleError(err)
	}

	stepper.StartTemporaryStepWithoutNewLine(
		"Opening Visual Studio Code",
	)

	env := output.Content.Env
	sshConfigHostKey := env.Name

	_, err = e.vscodeProcess.OpenOnRemote(
		sshConfigHostKey,
		constants.WorkspaceDirPath,
	)

	if err != nil {
		return handleError(err)
	}

	stepper.StopCurrentStep()

	e.presenter.PresentToView(EditResponse{
		Content: EditResponseContent{
			EnvName: env.Name,
		},
	})

	return nil
}
