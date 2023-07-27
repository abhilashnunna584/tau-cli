package application

import (
	"github.com/taubyte/tau-cli/cli/common"
	applicationI18n "github.com/taubyte/tau-cli/i18n/application"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	applicationPrompts "github.com/taubyte/tau-cli/prompts/application"
	"github.com/urfave/cli/v2"
)

func (link) Select() common.Command {
	return common.Create(
		&cli.Command{
			Action: _select,
		},
	)
}

func _select(ctx *cli.Context) error {
	application, deselect, err := applicationPrompts.GetSelectOrDeselect(ctx)
	if err != nil {
		return err
	}

	if deselect == true {
		err = applicationLib.Deselect(ctx, application.Name)
		if err != nil {
			return err
		}
		applicationI18n.Deselected(application.Name)
	} else {
		err = applicationLib.Select(ctx, application.Name)
		if err != nil {
			return err
		}
		applicationI18n.Selected(application.Name)
	}

	return nil
}
