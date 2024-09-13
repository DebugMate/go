package debugmate

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Occurrence struct {
	Payload url.Values
}

func OccurrenceFromEvent(event Event) (Occurrence, error) {
    traceJSON, err := json.Marshal(event.Trace)
    if err != nil {
        return Occurrence{}, err
    }

    requestJSON, err := json.Marshal(event.Request)
    if err != nil {
        return Occurrence{}, err
    }

    payload := url.Values{
        "exception": {event.Exception},
        "message":   {event.Message},
        "file":      {event.File},
        "type":      {event.Type},
        "url":       {event.URL},
        "code":      {strconv.Itoa(event.Code)},
        "trace":     {string(traceJSON)},
        "request":   {string(requestJSON)},
    }

    return Occurrence{Payload: payload}, nil
}
