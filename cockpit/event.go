package cockpit

import (
	"reflect"
	"runtime"
)

type Event struct {
	Exception string
	Message   string
	File      string
	Type      string
}

func EventFromError(err error) Event {
	_, file, _, _ := runtime.Caller(1)

	event := Event{
		Exception: reflect.TypeOf(err).String(),
		Message:   err.Error(),
		File:      file,
		Type:      "cli",
	}

	return event
}
