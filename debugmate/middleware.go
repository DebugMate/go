package debugmate

import (
	"fmt"
	"net/http"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}
				context := NewStackTraceContext()

				err := fmt.Errorf("%v", rvr)

				Catch(err, context.GetContext())

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))

				fmt.Println("Stack Trace Context:", context.GetContext())
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
