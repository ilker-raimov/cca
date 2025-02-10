package execute

import (
	"os"
	"os/exec"

	logger "github.com/sirupsen/logrus"
)

func Init() {
	cmd := exec.Command("go", "build", "-o", "docker_server", "docker/docker.go")

	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	logger.Info("Successfully built docker server.")
}
