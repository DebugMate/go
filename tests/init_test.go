package test

import (
	"testing"
	"github.com/DebugMate/go/debugmate"
)

func TestInit(t *testing.T) {
	options := debugmate.Options{
		Enabled: true,
		Token:   "test-token",
		Domain:  "http://example.com",
	}

	debugmate.Init(options)

	if !debugmate.Dbm.Options.Enabled {
		t.Errorf("Expected Enabled to be true, got %v", debugmate.Dbm.Options.Enabled)
	}
	if debugmate.Dbm.Options.Token != "test-token" {
		t.Errorf("Expected Token 'test-token', got %s", debugmate.Dbm.Options.Token)
	}
}