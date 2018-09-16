package command

import (
	"github.com/mitchellh/cli"
)

// RecordCommand initiates key strokes recording.
type RecordCommand struct {
	ShutdownCh <-chan struct{}
	UI         cli.Ui
}

var _ cli.Command = &RecordCommand{}

// Help method defines command instructions.
func (c *RecordCommand) Help() string {
	return `
Usage: keylogger record [options]

  ` + c.Synopsis() + `.

Options:
  -device=event0                   Device name.
  -log-path="/tmp/keylogger.log"   Path to the log file with key hits.
`
}

// Run method executes the command.
func (c *RecordCommand) Run(args []string) int {
	return 0
}

// Synopsis method provides short definition.
func (c *RecordCommand) Synopsis() string {
	return "Records any keys pressed on the selected device"
}
