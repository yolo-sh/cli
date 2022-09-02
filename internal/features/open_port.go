package features

import (
	"github.com/yolo-sh/yolo/features"
)

type OpenPortResponse struct {
	Error   error
	Content OpenPortResponseContent
}

type OpenPortResponseContent struct {
	EnvName            string
	EnvPublicIPAddress string
	PortOpened         string
	PortAlreadyOpened  bool
}

type OpenPortPresenter interface {
	PresentToView(OpenPortResponse)
}

type OpenPortOutputHandler struct {
	presenter OpenPortPresenter
}

func NewOpenPortOutputHandler(
	presenter OpenPortPresenter,
) OpenPortOutputHandler {

	return OpenPortOutputHandler{
		presenter: presenter,
	}
}

func (o OpenPortOutputHandler) HandleOutput(output features.OpenPortOutput) error {
	stepper := output.Stepper

	handleError := func(err error) error {
		stepper.StopCurrentStep()

		o.presenter.PresentToView(OpenPortResponse{
			Error: err,
		})

		return err
	}

	if output.Error != nil {
		return handleError(output.Error)
	}

	env := output.Content.Env

	stepper.StopCurrentStep()

	o.presenter.PresentToView(OpenPortResponse{
		Content: OpenPortResponseContent{
			EnvName:            env.Name,
			EnvPublicIPAddress: env.InstancePublicIPAddress,
			PortOpened:         output.Content.PortOpened,
			PortAlreadyOpened:  output.Content.PortAlreadyOpened,
		},
	})

	return nil
}
