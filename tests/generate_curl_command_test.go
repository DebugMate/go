package test

import (
	"net/http"
	"testing"
	"github.com/DebugMate/go/debugmate"
	"net/url"
)

func TestGenerateCurlCommand(t *testing.T) {
	req := &http.Request{
		Method: http.MethodPost,
		URL:    &url.URL{Path: "/test"},
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	curlCommand := debugmate.GenerateCurlCommand(req)
	expected := "curl -X POST '/test' -H 'Content-Type: application/json'"
	if curlCommand != expected {
		t.Errorf("expected %s, got %s", expected, curlCommand)
	}
}
