package code

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	codeDir      = "./code"       // Directory for Java files
	dockerImage  = "java_sandbox" // Docker image name
	containerCmd = "docker"       // Docker command
)

func Run(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	code, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, "Failed to read request", http.StatusBadRequest)

		return
	}

	if err := os.MkdirAll(codeDir, 0755); err != nil {
		http.Error(writer, "Failed to create code directory", http.StatusInternalServerError)

		return
	}

	filePath := filepath.Join(codeDir, "Main.java")

	if err := os.WriteFile(filePath, code, 0644); err != nil {
		http.Error(writer, "Failed to save Java file", http.StatusInternalServerError)

		return
	}

	cmd := exec.Command(containerCmd, "run", "--rm", "-v", fmt.Sprintf("%s:/app", codeDir), dockerImage)
	output, err := cmd.CombinedOutput()

	if err != nil {
		http.Error(writer, fmt.Sprintf("Execution error: %s\nOutput: %s", err.Error(), string(output)), http.StatusInternalServerError)

		return
	}

	writer.Write(output)
}
