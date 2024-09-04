package cockpit

import (
	"fmt"
	"net/http"
	"strings"
)

type Cockpit struct {
	Client  HTTPClient
	Options Options
}

var C Cockpit

func Init(options Options) {
	C = Cockpit{
		Client:  Client,
		Options: options,
	}
}

func Catch(err error) error {
	event := EventFromError(err)
	occurrence := OccurrenceFromEvent(event)

	return C.publish(occurrence)
}

func (c Cockpit) publish(occurrence Occurrence) error {
	if !c.Options.Enabled {
		return nil
	}

	body := strings.NewReader(occurrence.Payload.Encode())
	request, _ := http.NewRequest(http.MethodPost, c.endpoint(), body)

	request.Header.Set("X-COCKPIT-TOKEN", c.Options.Token)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := c.Client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("Ops, we're expecting an 201 status code, but received %d\n", response.StatusCode)
	}

	return nil
}

func (c Cockpit) endpoint() string {
	return strings.TrimSuffix(c.Options.Domain, "/") + "/webhook"
}
