package main

import (
	"github.com/pterm/pterm"
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/prompts"
	loginPrompts "github.com/taubyte/tau-cli/prompts/login"
	"github.com/urfave/cli/v2"
)

var WebTokenCommand = &cli.Command{
	Name: "token_from_web",
	Flags: []cli.Flag{
		flags.Provider,
	},
	Action: func(ctx *cli.Context) (err error) {
		var provider string
		if ctx.IsSet(flags.Provider.Name) == true {
			provider = ctx.String(flags.Provider.Name)
		}

		if len(provider) == 0 {
			provider, err = prompts.SelectInterface(loginPrompts.Providers, loginPrompts.GitProviderPrompt, loginPrompts.DefaultProvider)
			if err != nil {
				return err
			}
		}

		// New
		token, err := loginPrompts.TokenFromWeb(ctx, provider)
		if err != nil {
			return err
		}

		pterm.Success.Printfln("Got token `%s`", token)
		return nil
	},
}
