package controller

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"[T] \x1b[31;1m%s\t%s\t%s\t%s\x1b[0m",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
