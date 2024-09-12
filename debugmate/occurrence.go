package debugmate

import (
	"encoding/json"
	"net/url"
)

type Occurrence struct {
	Payload url.Values
}

// OccurrenceFromEvent cria uma instância de Occurrence a partir de um Event
func OccurrenceFromEvent(event Event) (Occurrence, error) {
	// Converte o slice de Trace para JSON
	traceJSON, err := json.Marshal(event.Trace)
	if err != nil {
		return Occurrence{}, err // Retorna um erro se a conversão falhar
	}

	// Inclui a string JSON no payload
	payload := url.Values{
		"exception": {event.Exception},
		"message":   {event.Message},
		"file":      {event.File},
		"type":      {event.Type},
		"trace":     {string(traceJSON)}, // Adiciona a string JSON ao payload
	}

	return Occurrence{Payload: payload}, nil
}
