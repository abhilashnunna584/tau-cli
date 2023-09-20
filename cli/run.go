package cli

import (
        argsLib "github.com/taubyte/tau-cli/cli/args"
        "github.com/taubyte/tau-cli/i18n"
)
// Run is the entry point for executing CLI commands.
// It accepts a variadic argument list of strings representing the command-line arguments.
// It creates a new CLI application, parses the command-line arguments, and runs the specified command.
func Run(args ...string) error {
        app, err := New()
        if err != nil {
                return i18n.AppCreateFailed(err)
        }

        if len(args) == 1 {
                return app.Run(args)
        }
// ParseArguments function is used to parse and prepare the command-line arguments.
        args = argsLib.ParseArguments(app.Flags, app.Commands, args...)

        return app.Run(args)
}

