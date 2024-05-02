package main

import (
	"fmt"
	"net/http"
	"time"
)

func InitializeHealthServer(cfg Config) *http.Server {
	mux := http.NewServeMux()

	healthCheckHandler := func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "available")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	mux.HandleFunc("/health", healthCheckHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Ports.Health),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
	http.DefaultTransport.(*http.Transport).IdleConnTimeout = 30 * time.Second

	return server
}
