package command

import (
	"flag"

	"github.com/mtojek/go-keylogger"

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
  -eventPath=/dev/input/event0     Event device path.
  -logPath="/tmp/keylogger.log"    Path to the log file with key hits.
`
}

// Run method executes the command.
func (c *RecordCommand) Run(args []string) int {
	var eventPath string
	var logPath string

	cmdFlags := flag.NewFlagSet("record", flag.ExitOnError)
	cmdFlags.StringVar(&eventPath, "eventPath", "/dev/input/event0", "Event device path")
	cmdFlags.StringVar(&logPath, "logPath", "/tmp/keylogger.log", "Path to the log file with key hits")
	err := cmdFlags.Parse(args)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output("Start recording...")

	var kr keylogger.KeyRecorder
	err = kr.Record(&keylogger.RecorderOptions{ShutdownCh: c.ShutdownCh, Logger: c.UI, EventPath: eventPath, LogPath: logPath})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output("Done.")
	return 0
}

// Synopsis method provides short definition.
func (c *RecordCommand) Synopsis() string {
	return "Records any keys pressed on the selected device"
}
