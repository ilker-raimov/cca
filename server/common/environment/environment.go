package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var variable_map map[string]string

func Init(filename string) {
	err := godotenv.Load(filename)

	if err != nil {
		panic(err)
	}

	variable_map = make(map[string]string)
}

func Get(key string) string {
	value, exists := variable_map[key]

	if !exists {
		value = os.Getenv(key)

		variable_map[key] = value
	}

	return value
}

func GetOrDefault(key string, fallback string) string {
	value := Get(key)

	if value == "" {
		return fallback
	}

	return value
}

func GetOrPanic(key string) string {
	value := Get(key)

	if value == "" {
		message := fmt.Sprintf("Missing environment variable: %s", key)

		panic(message)
	}

	return value
}
