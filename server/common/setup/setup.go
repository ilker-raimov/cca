package setup

import (
	"fmt"
	"strconv"
	"strings"
)

type Language int

const (
	JAVA_8 Language = iota
	JAVA_11
	JAVA_17
	JAVA_21
)

var stringToLanguage = map[string]Language{
	"JAVA_8":  JAVA_8,
	"JAVA_11": JAVA_11,
	"JAVA_17": JAVA_17,
	"JAVA_21": JAVA_21,
}

func (l Language) String() string {
	return [...]string{"JAVA_8", "JAVA_11", "JAVA_17", "JAVA_21"}[l]
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

func All() []string {
	return []string{JAVA_8.String(), JAVA_11.String(), JAVA_17.String(), JAVA_21.String()}
}

func Build(language Language, cmd string) string {
	var from string

	switch language {
	case JAVA_8:
		from = "openjdk:8"
	case JAVA_11:
		from = "openjdk:11"
	case JAVA_17:
		from = "openjdk:17"
	case JAVA_21:
		from = "openjdk:21"
	}

	return fmt.Sprintf("FROM %s\nCMD [\"./%s\"]", from, cmd)
}

func From(value string) Language {
	upper_value := strings.ToUpper(value)
	language := stringToLanguage[upper_value]

	return language
}
