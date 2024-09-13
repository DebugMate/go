package debugmate

import (
	"fmt"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status  int
	written bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.status = code
		rw.ResponseWriter.WriteHeader(code)
		rw.written = true
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}

				err := fmt.Errorf("Panic recovered: %v", rvr)
				Catch(err)

				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal Server Error"))
			}

			if rw.status >= 400 {
				err := fmt.Errorf("HTTP error: %d %s", rw.status, http.StatusText(rw.status))
				event, captureErr := EventFromErrorWithRequest(err, r)
				if captureErr != nil {
					fmt.Println("Error capturing request info:", captureErr)
				}

				occurrence, occErr := OccurrenceFromEvent(event)
				if occErr != nil {
					fmt.Println("Error creating occurrence from event:", occErr)
				}

				Dbm.publish(occurrence)
			}
		}()

		next.ServeHTTP(rw, r)
	})
}