package test

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/DebugMate/go/debugmate"
)

func TestEventFromError(t *testing.T) {
	err := errors.New("test error")
	stack := []debugmate.Trace{}

	event := debugmate.EventFromError(err, stack)

	if event.Exception != reflect.TypeOf(err).String() {
		t.Errorf("Expected exception %s, got %s", reflect.TypeOf(err).String(), event.Exception)
	}
	if event.Message != err.Error() {
		t.Errorf("Expected message %s, got %s", err.Error(), event.Message)
	}
	if event.File == "" {
		t.Error("Expected file to be non-empty")
	}
	if event.Type != "cli" {
		t.Errorf("Expected type 'cli', got %s", event.Type)
	}
	if len(event.Trace) != len(stack) {
		t.Errorf("Expected trace length %d, got %d", len(stack), len(event.Trace))
	}
}

func TestEventFromErrorWithRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	event, err := debugmate.EventFromErrorWithRequest(errors.New("test error"), req)
	if err != nil {
		t.Fatal(err)
	}

	if event.URL != req.URL.String() {
		t.Errorf("Expected URL %s, got %s", req.URL.String(), event.URL)
	}
	if event.Type != "web" {
		t.Errorf("Expected type 'web', got %s", event.Type)
	}
	if event.Request.Request.Method != req.Method {
		t.Errorf("Expected method %s, got %s", req.Method, event.Request.Request.Method)
	}
	if event.Request.Request.Curl == "" {
		t.Error("Expected curl command to be non-empty")
	}
	if event.Environment == nil {
		t.Error("Expected environment to be non-nil")
	}
}

func TestItCanGetAllRequiredValues(t *testing.T) {
	event := debugmate.EventFromError(errors.New("Some error"), debugmate.FormatStack())
	assert.Equal(t, "*errors.errorString", event.Exception)
	assert.Equal(t, "Some error", event.Message)
	assert.Contains(t, event.File, "event_test.go")
	assert.Equal(t, "cli", event.Type)
}