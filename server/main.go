package main

import (
	"fmt"
	"net/http"

	"github.com/ilker-raimov/cca/api/auth"
	"github.com/ilker-raimov/cca/interceptor"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&CustomFormatter{})
	logrus.Info("Server - starting")

	router := initRouter()

	http.ListenAndServe(":8080", router)

	logrus.Info("Server - stopped")
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s [%s]: %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)), nil
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(interceptor.LogInterceptor)

	router.Handle("/", http.FileServer(http.Dir("./browser")))

	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST")

	return router
}
