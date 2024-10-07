package utils

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadEnvOnce sync.Once

func LoadEnv() {
	loadEnvOnce.Do(func() {
		defaultEnv := `dev`
		env := os.Getenv(`APP_ENV`)

		if env == "" {
			os.Setenv(`APP_ENV`, defaultEnv)
			env = defaultEnv // default to development environment
		}

		godotenv.Load("config/.env." + env + ".local") // .env.<env>.local 1st precedence typically for secret environment specific variables
		godotenv.Load("config/.env." + env)            // .env.<env> 2nd precedence for environment based shared variables
		godotenv.Load("config/.env.local")             // .env.local 3rd precedence for secret variables for all environments
		godotenv.Load("config/.env")
	}) // finally 4th precedence for shared variables for all environments
}
