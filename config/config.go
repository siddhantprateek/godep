package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(ENV_VAR string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	envVariable := os.Getenv(ENV_VAR)

	// if envVariable == "" {
	// 	log.Fatalf("%s variable is not provided. CHECK .env file", ENV_VAR)
	// }

	return envVariable
}
