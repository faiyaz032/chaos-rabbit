package chaos

import (
	"net/http"
	"time"
)

func Latency(delay time.Duration) Middleware {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			time.Sleep(delay)

			next.ServeHTTP(w, r)

		})
	}
}
