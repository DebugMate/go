package debugmate

import (
	"reflect"
	"runtime"
)

type Event struct {
	Exception   string
	Message     string
	File        string
	Type        string
	Trace       string
}

func EventFromError(err error, stack string) Event {
	_, file, _, _ := runtime.Caller(1)

	event := Event{
		Exception: reflect.TypeOf(err).String(),
		Message:   err.Error(),
		File:      file,
		Type:      "cli",
		Trace:      stack,
	}

	return event
}
