package test

import (
	"bytes"
	"net/http"
	"testing"
	"github.com/DebugMate/go/debugmate"
)

func TestGetRequestBody(t *testing.T) {
	body := `{"key": "value"}`
	req, err := http.NewRequest("POST", "http://example.com", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	result := debugmate.GetRequestBody(req)
	if result["key"] != "value" {
		t.Errorf("Expected body key 'value', got %v", result["key"])
	}
}
