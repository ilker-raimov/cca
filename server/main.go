package main

import (
	"net/http"

	"github.com/ilker-raimov/cca/api/auth"
	"github.com/ilker-raimov/cca/interceptor"
	"github.com/ilker-raimov/cca/log"

	"github.com/gorilla/mux"
)

func main() {
	log.Init()

	log.Info("Server - starting")

	router := initRouter()

	http.ListenAndServe(":8080", router)

	log.Info("Server - stopped")
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(interceptor.LogInterceptor)

	router.Handle("/", http.FileServer(http.Dir("./browser")))

	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST")

	return router
}
