package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type DevicesCommand struct {
	Ui cli.Ui
}

var _ cli.Command = &DevicesCommand{}

func (c *DevicesCommand) Help() string {
	helpText := `
Usage: keylogger devices [options]

  ` + c.Synopsis() + `.
`
	return strings.TrimSpace(helpText)
}

func (c *DevicesCommand) Run(args []string) int {
	return 0
}

func (c *DevicesCommand) Synopsis() string {
	return "Lists available input devices"
}
