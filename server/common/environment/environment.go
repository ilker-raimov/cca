package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init(filename string) {
	err := godotenv.Load(filename)

	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
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
