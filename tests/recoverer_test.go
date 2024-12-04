package test

import (
	"github.com/DebugMate/go/debugmate"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoverer(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "status 500",
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "status 400",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			handler := debugmate.Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, rr.Code)
			}
		})
	}
}
