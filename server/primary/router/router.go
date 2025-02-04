package router

import (
	"net/http"

	"github.com/ilker-raimov/cca/common/interceptor"
	"github.com/ilker-raimov/cca/primary/api/auth"
	"github.com/ilker-raimov/cca/primary/api/code"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.Use(interceptor.LogInterceptor)

	router.Handle("/", http.FileServer(http.Dir("./browser")))

	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/auth/register", auth.Register).Methods("POST")
	router.HandleFunc("/api/code/run", code.Run).Methods("POST")

	return router
}
