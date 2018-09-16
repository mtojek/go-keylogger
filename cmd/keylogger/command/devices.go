package command

import (
	"bytes"
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/mtojek/go-keylogger"
)

// DevicesCommand uses DeviceManager to list available devices.
type DevicesCommand struct {
	UI cli.Ui
}

var _ cli.Command = &DevicesCommand{}

// Help method defines command instructions.
func (c *DevicesCommand) Help() string {
	return `
Usage: keylogger devices

  ` + c.Synopsis() + `.
`
}

// Run method runs the command.
func (c *DevicesCommand) Run(args []string) int {
	var dm keylogger.DeviceManager
	devices, err := dm.ListDevices()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	if len(devices) == 0 {
		c.UI.Output("No event devices available.")
		return 0
	}

	var listing bytes.Buffer
	listing.WriteString("Available event devices:\n")
	for _, device := range devices {
		listing.WriteString(fmt.Sprintf("  %s (name: \"%s\", path: %s)\n", device.ID, device.Name,
			device.EventPath))
	}

	c.UI.Output(listing.String())
	return 0
}

// Synopsis method provides short definition.
func (c *DevicesCommand) Synopsis() string {
	return "Lists available input devices"
}
