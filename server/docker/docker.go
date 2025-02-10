package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilker-raimov/cca/common/interceptor"
	"github.com/ilker-raimov/cca/common/log"
	"github.com/ilker-raimov/cca/docker/api/run"
	"github.com/ilker-raimov/cca/docker/code"
	logger "github.com/sirupsen/logrus"
)

func main() {
	log.Init()
	code.Init()

	stop := make(chan int)
	router := mux.NewRouter()

	router.Use(interceptor.LogInterceptor)

	router.HandleFunc("/api/run/compile", run.Compile).Methods("POST")
	router.HandleFunc("/api/run/check", run.Check).Methods("POST")
	router.HandleFunc("/api/run/test", run.Test).Methods("POST")
	router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		stop <- 1
	}).Methods("GET")

	logger.Info("Server - starting")

	server := &http.Server{Addr: "localhost:8080", Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Info("Server stopped.")
			} else {
				logger.Errorf("Server did not start due to: %s", err)
			}
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
