package debugmate

import (
	"reflect"
	"runtime"
)

type Event struct {
	Exception string
	Message   string
	File      string
	Type      string
	Trace     []Trace
}

// EventFromError cria um evento a partir de um erro e do trace
func EventFromError(err error, stack []Trace) Event {
	_, file, _, _ := runtime.Caller(1)

	event := Event{
		Exception: reflect.TypeOf(err).String(),
		Message:   err.Error(),
		File:      file,
		Type:      "cli",
		Trace:     stack,
	}

	return event
}
