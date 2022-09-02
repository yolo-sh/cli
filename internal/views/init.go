package views

import "github.com/yolo-sh/cli/internal/presenters"

type InitView struct {
	BaseView
}

func NewInitView(baseView BaseView) InitView {
	return InitView{
		BaseView: baseView,
	}
}

func (i InitView) View(data presenters.InitViewData) {
	if data.Error == nil {
		if data.Content.ShowAsWarning {
			i.ShowWarningView(
				data.Content.Message,
				data.Content.Subtext,
			)
			return
		}

		i.ShowSuccessView(
			data.Content.Message,
			data.Content.Subtext,
		)
		return
	}

	i.ShowErrorView(data.Error)
}
