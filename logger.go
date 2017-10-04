package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// NewLogger constructs a logger.
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "[GoRestAPI] ", 0)
}

// Log wraps up information about each request.
func Log(logger *log.Logger, h http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, req)
			logger.Printf("%s\t%s\t%s\t%s",
				req.Method, req.RequestURI, name, time.Since(start),
			)
		})
}
