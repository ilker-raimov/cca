package router

import (
	"net/http"

	"github.com/ilker-raimov/cca/common/interceptor"
	"github.com/ilker-raimov/cca/primary/api/auth"
	"github.com/ilker-raimov/cca/primary/api/code"
	"github.com/ilker-raimov/cca/primary/api/competition"
	"github.com/ilker-raimov/cca/primary/api/task"
	"github.com/ilker-raimov/cca/primary/api/users"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	browser := http.Dir("browser")
	fs := http.FileServer(browser)

	router.Use(interceptor.LogInterceptor)

	router.Handle("/", fs)
	router.Handle("/main.js", fs)
	router.Handle("/main.css", fs)

	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/auth/register", auth.Register).Methods("POST")

	router.HandleFunc("/api/competitions/languages", competition.Languages).Methods("GET")
	router.HandleFunc("/api/competitions", competition.List).Methods("GET")
	router.HandleFunc("/api/competitions", competition.Create).Methods("POST")
	router.HandleFunc("/api/competitions/{competition_id}", competition.Get).Methods("GET")
	router.HandleFunc("/api/competitions/{competition_id}/tasks", task.List).Methods("GET")
	router.HandleFunc("/api/competitions/{competition_id}/tasks", task.Create).Methods("POST")
	router.HandleFunc("/api/competitions/{competition_id}/tasks/{task_id}", task.Get).Methods("GET")

	router.HandleFunc("/api/users/roles", users.Roles).Methods("GET")

	router.HandleFunc("/api/code/run", code.Run).Methods("POST")

	return router
}
