package keylogger

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
)

const (
	logFileActualPath   = "test-files/log_file_actual"
	logFileExpectedPath = "test-files/log_file_expected"
)

type mockedLogger struct{}

func (ml *mockedLogger) Error(message string) {}

func TestKeyRecorder_Record(t *testing.T) {
	// given
	options := &RecorderOptions{
		Logger: &mockedLogger{},

		EventPath: "test-files/keyboard_dump",
		LogPath:   logFileActualPath,
	}
	defer os.Remove(logFileActualPath)

	// when
	var kr KeyRecorder
	err := kr.doRecording(options)
	if errors.Cause(err) != io.EOF {
		t.Errorf("reading events from ordinary files should end up with EOF, actual: %v", err)
	}

	// then
	actualFile, err := ioutil.ReadFile(logFileActualPath)
	if err != nil {
		t.Errorf("can't read actual log file: %v", err)
	}

	expectedFile, err := ioutil.ReadFile(logFileExpectedPath)
	if err != nil {
		t.Errorf("can't read expected log file: %v", err)
	}

	if string(expectedFile) != string(actualFile) {
		t.Errorf("expected and actual files are different.\nExpected: %v\nActual:   %v", expectedFile, actualFile)
	}
}
