package setup

import (
	"fmt"
	"strconv"
	"strings"
)

//GOOS=linux GOARCH=amd64 go build -o myapp main.go

type Language int

const (
	JAVA_17 Language = iota
	JAVA_21
	GO_120
	CPP_14
	CPP_20
)

func (l Language) String() string {
	return [...]string{"JAVA_17", "JAVA_21", "CPP_9", "CPP_11"}[l]
}

func (l Language) GetName() string {
	versioned_name := l.String()

	return strings.Split(versioned_name, "_")[0]
}

func (l Language) GetVersion() (int, error) {
	versioned_name := l.String()
	version := strings.Split(versioned_name, "_")[1]

	return strconv.Atoi(version)
}

func Build(language Language, cmd string) string {
	var from string

	switch language {
	case JAVA_17:
		from = "openjdk:17"
	case JAVA_21:
		from = "openjdk:21"
	case GO_120:
		from = "golang:1.20"
	case CPP_14:
		from = "gcc:14"
	case CPP_20:
		from = "gcc:17"
	}

	return fmt.Sprintf("FROM %s\nCMD [\"./%s\"]", from, cmd)
}
