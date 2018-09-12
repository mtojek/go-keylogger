package main

import (
	"os"
	"os/signal"

	"github.com/mitchellh/cli"
	"github.com/mtojek/go-keylogger/cmd/keylogger/command"
)

// Commands is the mapping of all the available commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"devices": func() (cli.Command, error) {
			return &command.DevicesCommand{
				Ui: ui,
			}, nil
		},
		"record": func() (cli.Command, error) {
			return &command.RecordCommand{
				ShutdownCh: makeShutdownCh(),
				Ui:         ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Ui:      ui,
				Version: Version,
			}, nil
		},
	}
}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}
