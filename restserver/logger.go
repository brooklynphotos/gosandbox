package main

import (
	"log"
	"net/http"
	"time"
)

/**
 * Logger decorates handlers with loggers
 */
func Logger(name string, inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, req)
		log.Printf("%s\t%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			name,
			time.Since(start),
		)
	})
}
