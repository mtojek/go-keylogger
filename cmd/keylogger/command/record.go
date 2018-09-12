package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type RecordCommand struct {
	ShutdownCh <-chan struct{}
	Ui         cli.Ui
}

var _ cli.Command = &RecordCommand{}

func (c *RecordCommand) Help() string {
	helpText := `
Usage: keylogger record [options]

  ` + c.Synopsis() + `.

Options:
  -device=event0                   Device name.
  -log-path="/tmp/keylogger.log"   Path to the log file with key hits.
`
	return strings.TrimSpace(helpText)
}

func (c *RecordCommand) Run(args []string) int {
	return 0
}

func (c *RecordCommand) Synopsis() string {
	return "Records any keys pressed on the selected device"
}
