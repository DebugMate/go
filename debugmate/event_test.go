package debugmate

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"runtime/debug"
	"testing"
)

func TestItCanGetAllRequiredValues(t *testing.T) {
	event := EventFromError(errors.New("Some error"), string(debug.Stack()))

	assert.Equal(t, "*errors.errorString", event.Exception)
	assert.Equal(t, "Some error", event.Message)
	assert.Contains(t, event.File, "event_test.go")
	assert.Equal(t, "cli", event.Type)
}
