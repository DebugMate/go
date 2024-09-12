package debugmate

import (
	"fmt"
	"net/http"
	"strings"
)

type DebugMate struct {
	Client  HTTPClient
	Options Options
}

var Dbm DebugMate

func Init(options Options) {
	Dbm = DebugMate{
		Client:  Client,
		Options: options,
	}
}

func Catch(err error) error {
	stackTraceContext := NewStackTraceContext() // Captura o trace automaticamente
	event := EventFromError(err, stackTraceContext.GetContext())
	occurrence, err := OccurrenceFromEvent(event)
	if err != nil {
		return err
	}

	return Dbm.publish(occurrence)
}

func (dbm DebugMate) publish(occurrence Occurrence) error {
	if !dbm.Options.Enabled {
		return nil
	}

	body := strings.NewReader(occurrence.Payload.Encode())
	request, _ := http.NewRequest(http.MethodPost, dbm.endpoint(), body)

	request.Header.Set("X-DEBUGMATE-TOKEN", dbm.Options.Token)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := dbm.Client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("ops, we're expecting a 201 status code, but received %d", response.StatusCode)
	}

	return nil
}

func (dbm DebugMate) endpoint() string {
	return strings.TrimSuffix(dbm.Options.Domain, "/") + "/api/capture"
}
