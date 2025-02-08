package environment

import (
	"os"

	env "github.com/joho/godotenv"
)

func Init(filename string) {
	env.Load(filename)
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
