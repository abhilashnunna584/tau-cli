package projectTable

import (
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts"
	projectPrompts "github.com/taubyte/tau-cli/prompts/project"
	"github.com/urfave/cli/v2"
)

func Confirm(ctx *cli.Context, project *projectLib.Project, prompt string) bool {
	var visibilityString string
	if project.Public == true {
		visibilityString = projectPrompts.Public
	} else {
		visibilityString = projectPrompts.Private
	}

	return prompts.ConfirmData(ctx, prompt, [][]string{
		{"Name", project.Name},
		{"Description", project.Description},
		{"Visibility", visibilityString},
	})
}
