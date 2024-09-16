package test

import (
	"errors"
	"testing"
	"net/http"
	"github.com/DebugMate/go/debugmate"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestCatch(t *testing.T) {
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "http://example.com/api/capture" {
				t.Errorf("Expected request URL 'http://example.com/api/capture', got %s", req.URL.String())
			}
			return &http.Response{
				StatusCode: http.StatusCreated,
			}, nil
		},
	}

	options := debugmate.Options{
		Enabled: true,
		Token:   "test-token",
		Domain:  "http://example.com",
	}

	debugmate.Init(options)
	debugmate.Dbm.Client = mockClient

	err := debugmate.Catch(errors.New("test error"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}