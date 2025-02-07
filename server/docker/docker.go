package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilker-raimov/cca/common/interceptor"
	"github.com/ilker-raimov/cca/common/log"
	"github.com/ilker-raimov/cca/docker/api/run"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Init()

	router := mux.NewRouter()

	router.Use(interceptor.LogInterceptor)

	router.HandleFunc("/api/run/compile", run.Compile).Methods("POST")

	logrus.Info("Server - starting")

	http.ListenAndServe("localhost:8080", router)

	logrus.Info("Server - stopped")
}
