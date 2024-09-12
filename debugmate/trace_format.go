package debugmate

import (
	"runtime"
	"strings"
)

type StackTraceContext struct {
	Stack []Trace
}

func NewStackTraceContext() *StackTraceContext {
	return &StackTraceContext{
		Stack: formatStack(),
	}
}

func formatStack() []Trace {
	var traces []Trace
	stack := make([]byte, 1024)
	n := runtime.Stack(stack, false)
	stackStr := string(stack[:n])

	lines := strings.Split(stackStr, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "\t") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				trace := Trace{
					File:     parts[0],
					Line:     0,
					Function: strings.Join(parts[1:], " "),
					Preview:  []string{line},
				}
				traces = append(traces, trace)
			}
		}
	}

	return traces
}

type Trace struct {
	File     string   `json:"file"`
	Line     int      `json:"line"`
	Function string   `json:"function"`
	Preview  []string `json:"preview"`
}

func (s *StackTraceContext) GetContext() []Trace {
	return s.Stack
}
