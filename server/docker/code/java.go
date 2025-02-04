package code

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/ilker-raimov/cca/common/util/file"
	"github.com/sirupsen/logrus"
)

const (
	JAVA_FILE         = "Main.java"
	COULD_NOT_COMPILE = "Could not compile due to error. Retry..."
)

type Java struct {
}

func New() *Java {
	return &Java{}
}

func (j Java) Compile(code []byte) (bool, string, error) {
	logrus.Infof("Compile: %s", string(code))

	id := uuid.New().String()

	logrus.Infof("ID: %s", id)

	if err := os.MkdirAll(id, 0755); err != nil {
		logrus.Warn("Could not create dir")

		return false, COULD_NOT_COMPILE, err
	}

	input_path := filepath.Join(id, JAVA_FILE)

	if err := os.WriteFile(input_path, code, 0644); err != nil {
		logrus.Warn("Could not create file")

		return false, COULD_NOT_COMPILE, err
	}

	logrus.Infof("Executing: javac -d %s %s", id, input_path)

	cmd := exec.Command("javac", "-d", id, input_path)
	output, err := cmd.CombinedOutput()

	if err != nil {
		logrus.Warn("Could not execute command")

		return false, COULD_NOT_COMPILE, err
	}

	logrus.Info("Checking for successful compilation")

	result_path := filepath.Join(id, "Main.class")
	ok := file.Exists(result_path)

	var msg string

	if ok {
		msg = "Successfully compiled"
	} else {
		msg = string(output)
	}

	logrus.Info(msg)
	logrus.Info("Cleaning up")

	if err := os.RemoveAll(id); err != nil {
		logrus.Warn("Could not cleanup")

		return ok, msg, err
	}

	logrus.Info("Successful cleanup")

	return ok, msg, nil
}
