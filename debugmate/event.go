package debugmate

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"
	"os"
)

type Request struct {
	Request struct {
		URL  string `json:"url"`
		Method string `json:"method"`
		Curl   string `json:"curl"`
	} `json:"request"`
	Headers      map[string][]string `json:"headers"`
	QueryString  map[string][]string `json:"query_string"`
	Body         map[string]interface{} `json:"body"`
	Cookies      []Cookie             `json:"cookies"`
}

type Environment struct {
	Group     string            `json:"group"`
	Variables map[string]string `json:"variables"`
}

type Cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Event struct {
	Exception string   `json:"exception"`
	Message   string   `json:"message"`
	File      string   `json:"file"`
	Type      string   `json:"type"`
	Code      int      `json:"code"`
	URL       string   `json:"url"`
	Trace     []Trace  `json:"trace"`
	Request   Request  `json:"request"`
	Environment []Environment `json:"environment"`
}

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

func EventFromErrorWithRequest(err error, r *http.Request) (Event, error) {
	stack := NewStackTraceContext().GetContext()

	event := EventFromError(err, stack)
	event.Type = "web"
	event.URL = r.URL.String()
	event.Code = http.StatusInternalServerError

	event.Request = Request{
		Request: struct {
			URL    string `json:"url"`
			Method string `json:"method"`
			Curl   string `json:"curl"`
		}{
			URL:    r.URL.String(),
			Method: r.Method,
			Curl:   generateCurlCommand(r),
		},
		Headers:     r.Header,
		QueryString: r.URL.Query(),
		Body:        getRequestBody(r),
		Cookies:     getCookies(r),
	}

	event.Environment = getEnvironment()

	return event, nil
}

func getRequestBody(r *http.Request) map[string]interface{} {
	var body map[string]interface{}
	if r.Body != nil {
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		json.Unmarshal(buf.Bytes(), &body)
	}
	return body
}

func getCookies(r *http.Request) []Cookie {
	var cookies []Cookie
	for _, cookie := range r.Cookies() {
		cookies = append(cookies, Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
	return cookies
}

func generateCurlCommand(r *http.Request) string {
	curl := "curl -X " + r.Method + " '" + r.URL.String() + "'"

	for header, values := range r.Header {
		for _, value := range values {
			curl += " -H '" + header + ": " + value + "'"
		}
	}

	body := getRequestBody(r)
	if len(body) > 0 {
		jsonBody, _ := json.Marshal(body)
		curl += " -d '" + string(jsonBody) + "'"
	}

	return curl
}

func (e *Event) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}


func getEnvironment() []Environment {
	return []Environment{
		{
			Group: "System",
			Variables: map[string]string{
				"Go Version":     runtime.Version(),
				"OS":             runtime.GOOS,
				"Architecture":   runtime.GOARCH,
				"CPU Cores":      strconv.Itoa(runtime.NumCPU()),
				"Go Root":        runtime.GOROOT(),
				"Go Path":        os.Getenv("GOPATH"),
				"Current Dir":    getCurrentDir(),
				"Environment Date Time": time.Now().Format(time.RFC3339),
			},
		},
	}
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return dir
}