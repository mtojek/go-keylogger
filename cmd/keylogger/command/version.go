package command

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type VersionCommand struct {
	Ui      cli.Ui
	Version string
}

var _ cli.Command = &VersionCommand{}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Run(args []string) int {
	c.Ui.Output(fmt.Sprintf("keylogger v%s", c.Version))
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Prints the application version"
}
