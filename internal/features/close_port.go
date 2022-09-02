package features

import (
	"github.com/yolo-sh/yolo/features"
)

type ClosePortResponse struct {
	Error   error
	Content ClosePortResponseContent
}

type ClosePortResponseContent struct {
	EnvName           string
	PortClosed        string
	PortAlreadyClosed bool
}

type ClosePortPresenter interface {
	PresentToView(ClosePortResponse)
}

type ClosePortOutputHandler struct {
	presenter ClosePortPresenter
}

func NewClosePortOutputHandler(
	presenter ClosePortPresenter,
) ClosePortOutputHandler {

	return ClosePortOutputHandler{
		presenter: presenter,
	}
}

func (c ClosePortOutputHandler) HandleOutput(output features.ClosePortOutput) error {
	stepper := output.Stepper

	handleError := func(err error) error {
		stepper.StopCurrentStep()

		c.presenter.PresentToView(ClosePortResponse{
			Error: err,
		})

		return err
	}

	if output.Error != nil {
		return handleError(output.Error)
	}

	env := output.Content.Env

	stepper.StopCurrentStep()

	c.presenter.PresentToView(ClosePortResponse{
		Content: ClosePortResponseContent{
			EnvName:           env.Name,
			PortClosed:        output.Content.PortClosed,
			PortAlreadyClosed: output.Content.PortAlreadyClosed,
		},
	})

	return nil
}
