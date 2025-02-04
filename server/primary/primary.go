package primary

import (
	"net/http"

	"github.com/ilker-raimov/cca/common/log"
	"github.com/ilker-raimov/cca/primary/router"
)

func Start() {
	log.Init()

	router := router.Init()

	log.Info("Server - starting")

	http.ListenAndServe("localhost:8080", router)

	log.Info("Server - stopped")
}
