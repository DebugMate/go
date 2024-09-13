package debugmate

import (
	"encoding/json"
	"net/url"
)

type Occurrence struct {
	Payload url.Values
}

func OccurrenceFromEvent(event Event) (Occurrence, error) {
	traceJSON, err := json.Marshal(event.Trace)
	if err != nil {
		return Occurrence{}, err
	}

	payload := url.Values{
		"exception": {event.Exception},
		"message":   {event.Message},
		"file":      {event.File},
		"type":      {event.Type},
		"trace":     {string(traceJSON)},
	}

	return Occurrence{Payload: payload}, nil
}
