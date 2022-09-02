package features

import (
	"os"

	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/features"
)

type UninstallResponse struct {
	Error   error
	Content UninstallResponseContent
}

type UninstallResponseContent struct {
	YoloAlreadyUninstalled    bool
	SuccessMessage            string
	AlreadyUninstalledMessage string
	YoloExecutablePath        string
	YoloConfigDirPath         string
}

type UninstallPresenter interface {
	PresentToView(UninstallResponse)
}

type UninstallOutputHandler struct {
	presenter UninstallPresenter
}

func NewUninstallOutputHandler(
	presenter UninstallPresenter,
) UninstallOutputHandler {

	return UninstallOutputHandler{
		presenter: presenter,
	}
}

func (u UninstallOutputHandler) HandleOutput(output features.UninstallOutput) error {
	output.Stepper.StopCurrentStep()

	handleError := func(err error) error {
		u.presenter.PresentToView(UninstallResponse{
			Error: err,
		})

		return err
	}

	if output.Error != nil {
		return handleError(output.Error)
	}

	yoloExecutablePath, err := os.Executable()

	if err != nil {
		return handleError(err)
	}

	yoloConfigDirPath := system.UserConfigDir()

	u.presenter.PresentToView(UninstallResponse{
		Content: UninstallResponseContent{
			YoloAlreadyUninstalled:    output.Content.YoloAlreadyUninstalled,
			SuccessMessage:            output.Content.SuccessMessage,
			AlreadyUninstalledMessage: output.Content.AlreadyUninstalledMessage,
			YoloExecutablePath:        yoloExecutablePath,
			YoloConfigDirPath:         yoloConfigDirPath,
		},
	})

	return nil
}
