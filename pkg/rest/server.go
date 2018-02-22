package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func startHTTPService() error {
	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
	return nil
}
