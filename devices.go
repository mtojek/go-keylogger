package keylogger

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// Device describes generic input interface.
type Device struct {
	EventPath string
	ID        string
	Name      string
}

// DeviceManager allows for listing event devices.
type DeviceManager struct{}

var eventDevice = regexp.MustCompile("event[0-9]+")

// ListDevices provides a set of available event devices.
func (dm *DeviceManager) ListDevices() ([]Device, error) {
	// Ensure root permissions
	root := new(Root)
	err := root.Ensure()
	if err != nil {
		return nil, errors.Wrap(err, "can't list devices")
	}

	// Read all items in the directory
	inputs, err := ioutil.ReadDir("/sys/class/input")
	if err != nil {
		return nil, errors.Wrap(err, "can't read directory with input devices")
	}

	// Filter event devices
	var filtered []Device
	for _, input := range inputs {
		deviceID := input.Name()

		if !eventDevice.MatchString(deviceID) {
			continue
		}

		eventPath := filepath.Join("/dev/input", deviceID)
		_, err := os.Stat(eventPath)
		if err != nil {
			return nil, errors.Wrap(err, "can't read input directory")
		}

		dn, err := ioutil.ReadFile(filepath.Join("/sys/class/input", deviceID, "device/name"))
		if err != nil {
			return nil, errors.Wrap(err, "can't read device name")
		}
		deviceName := strings.TrimSpace(string(dn))

		filtered = append(filtered, Device{Name: deviceName, ID: deviceID, EventPath: eventPath})
	}
	return filtered, nil
}
