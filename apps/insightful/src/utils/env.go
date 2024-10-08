package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	Port              string
	AWSRegion         string
	AWSProfile        string
	CognitoUserPoolId string
}

var env ENV

func init() {
	defaultEnv := `dev`
	appenv := os.Getenv(`APP_ENV`)

	if appenv == "" {
		os.Setenv(`APP_ENV`, defaultEnv)
		appenv = defaultEnv // default to development environment
	}

	godotenv.Load("config/.env." + appenv + ".local") // .env.<env>.local 1st precedence typically for secret environment specific variables
	godotenv.Load("config/.env." + appenv)            // .env.<env> 2nd precedence for environment based shared variables
	godotenv.Load("config/.env.local")                // .env.local 3rd precedence for secret variables for all environments
	godotenv.Load("config/.env")

	env = ENV{
		Port:              os.Getenv("PORT"),
		AWSRegion:         os.Getenv("AWS_REGION"),
		AWSProfile:        os.Getenv("AWS_PROFILE"),
		CognitoUserPoolId: os.Getenv("COGNITO_USER_POOL_ID"),
	}
}

func LoadEnv() ENV {
	return env
}
