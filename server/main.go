package main

import (
	"net/http"

	"github.com/ilker-raimov/cca/api/auth"
	"github.com/ilker-raimov/cca/interceptor"
	"github.com/ilker-raimov/cca/log"

	"github.com/gorilla/mux"

	logger "github.com/sirupsen/logrus"
)

func main() {
	log.Init()

	logger.Info("Server - starting")

	router := initRouter()

	http.ListenAndServe(":8080", router)

	logger.Info("Server - stopped")
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(interceptor.LogInterceptor)

	router.Handle("/", http.FileServer(http.Dir("./browser")))

	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/auth/register", auth.Register).Methods("POST")

	return router
}
