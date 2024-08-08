package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("NO .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HHTP_PORT", ":8080"))

	return config
}

func coalesce(key string, defaultVaule interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultVaule
}
