package test

import (
	"runtime"
	"github.com/DebugMate/go/debugmate"
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	env := debugmate.GetEnvironment()

	if len(env) != 1 {
		t.Errorf("expected 1 environment group, got %d", len(env))
	}

	systemEnv := env[0]
	if systemEnv.Group != "System" {
		t.Errorf("expected group 'System', got %s", systemEnv.Group)
	}

	if systemEnv.Variables["Go Version"] != runtime.Version() {
		t.Errorf("expected Go Version %s, got %s", runtime.Version(), systemEnv.Variables["Go Version"])
	}
}
