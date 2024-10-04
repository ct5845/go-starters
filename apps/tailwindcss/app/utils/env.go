package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var defaultEnv string = "development"

func Env(key string) {
	if key == "" {
		key = "APP_ENV"
	}

	env := os.Getenv(key)

	if env == "" {
		fmt.Printf("%sENV[%s] not set, defaulting to: %s%s\n", Colors["Yellow"], key, Colors["Reset"], defaultEnv)
		os.Setenv(key, defaultEnv)
		env = defaultEnv // default to development environment
	}

	godotenv.Load(".env." + env + ".local") // .env.<env>.local 1st precedence typically for secret environment specific variables
	godotenv.Load(".env." + env)            // .env.<env> 2nd precedence for environment based shared variables
	godotenv.Load(".env.local")             // .env.local 3rd precedence for secret variables for all environments
	godotenv.Load()                         // finally 4th precedence for shared variables for all environments
}
