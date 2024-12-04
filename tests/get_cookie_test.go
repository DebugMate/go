package test

import (
	"github.com/DebugMate/go/debugmate"
	"net/http"
	"testing"
)

func TestGetCookies(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{Name: "session", Value: "12345"}
	req.AddCookie(cookie)

	cookies := debugmate.GetCookies(req)
	if len(cookies) != 1 {
		t.Fatalf("Expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "session" || cookies[0].Value != "12345" {
		t.Errorf("Expected cookie 'session' with value '12345', got %v", cookies[0])
	}
}
