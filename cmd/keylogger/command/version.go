package command

import (
	"fmt"

	"github.com/mitchellh/cli"
)

// VersionCommand prints application version.
type VersionCommand struct {
	UI      cli.Ui
	Version string
}

var _ cli.Command = &VersionCommand{}

// Help method defines command instructions.
func (c *VersionCommand) Help() string {
	return `
Usage: keylogger version

  ` + c.Synopsis() + `.
`
}

// Run method executes the command.
func (c *VersionCommand) Run(args []string) int {
	c.UI.Output(fmt.Sprintf("keylogger v%s", c.Version))
	return 0
}

// Synopsis method provides short definition.
func (c *VersionCommand) Synopsis() string {
	return "Prints the application version"
}
