package debugmate

import "net/url"

type Occurrence struct {
	Payload url.Values
}

func OccurrenceFromEvent(event Event) Occurrence {
	payload := url.Values{
		"exception":   {event.Exception},
		"message":     {event.Message},
		"file":        {event.File},
		"type":        {event.Type},
		"trace":       {event.Trace},
	}

	return Occurrence{payload}
}
