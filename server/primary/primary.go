package primary

import (
	"net/http"

	"github.com/ilker-raimov/cca/common/environment"
	"github.com/ilker-raimov/cca/common/log"
	"github.com/ilker-raimov/cca/primary/execute"
	"github.com/ilker-raimov/cca/primary/router"

	"github.com/sirupsen/logrus"
)

func Start() {
	log.Init()
	execute.Init()
	environment.Init("primary.env")

	router := router.Init()

	logrus.Info("Server - starting")

	http.ListenAndServe("localhost:8080", router)

	logrus.Info("Server - stopped")
}
