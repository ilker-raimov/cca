package code

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	JAVA_FILE         = "Main.java"
	COULD_NOT_COMPILE = "Could not compile due to error. Retry..."
)

type Java struct {
	Version int
}

func NewJava(version int) *Java {
	return &Java{
		Version: version,
	}
}

func (j *Java) Compile(code string) (bool, string, error) {
	logrus.Infof("Compile: %s", code)

	id := uuid.New().String()

	logrus.Infof("ID: %s", id)

	if err := os.MkdirAll(id, 0755); err != nil {
		logrus.Warn("Could not create dir")

		return false, COULD_NOT_COMPILE, err
	}

	input_path := filepath.Join(id, JAVA_FILE)

	if err := os.WriteFile(input_path, []byte(code), 0644); err != nil {
		logrus.Warn("Could not create file")

		j.Cleanup(id)

		return false, COULD_NOT_COMPILE, err
	}

	logrus.Infof("Executing: javac -d %s %s", id, input_path)

	cmd := exec.Command("javac", "-d", id, input_path)
	output, err := cmd.CombinedOutput()

	if err != nil {
		logrus.Warn("Could not execute command")

		to_replace := fmt.Sprintf("%s/Main.java", id)
		sanitized_output := strings.Replace(string(output), to_replace, "line", -1)

		j.Cleanup(id)

		return false, sanitized_output, err
	}

	j.Cleanup(id)

	return true, "Successfully compiled", nil
}

func (j *Java) Test(code string, tests []any) {

}

func (j *Java) Cleanup(id string) {
	logrus.Info("Cleaning up")

	if err := os.RemoveAll(id); err != nil {
		logrus.Warn("Could not cleanup")

		return
	}

	logrus.Info("Successful cleanup")
}
