package debugmate

import (
	"runtime"
	"strconv"
	"strings"
)

type Trace struct {
	File     string   `json:"file"`
	Line     int      `json:"line"`
	Function string   `json:"function"`
	Preview  []string `json:"preview"`
}

// StackTraceContext armazena o contexto do stack trace
type StackTraceContext struct {
	Stack []Trace
}

// NewStackTraceContext cria uma nova instância de StackTraceContext a partir de um panic
func NewStackTraceContext() *StackTraceContext {
	return &StackTraceContext{
		Stack: formatStack(),
	}
}

// formatStack captura e formata o stack trace
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
					File:     strings.Split(parts[0], ":")[0],
					Line:     parseLine(parts[0]),
					Function: strings.Join(parts[1:], " "),
					Preview:  []string{line}, // Adiciona a linha completa como preview
				}
				traces = append(traces, trace)
			}
		}
	}

	return traces
}

// parseLine extrai o número da linha do formato "file:line"
func parseLine(fileLine string) int {
	parts := strings.Split(fileLine, ":")
	if len(parts) > 1 {
		line, err := strconv.Atoi(parts[1])
		if err == nil {
			return line
		}
	}
	return 0
}

// GetContext retorna o contexto do stack trace
func (s *StackTraceContext) GetContext() []Trace {
	return s.Stack
}
