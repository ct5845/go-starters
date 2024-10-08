package cognito

import (
	"log"
	"open_api/src/utils"
)

type Config struct {
	Region     string
	Profile    string
	UserPoolId string
}

var config Config

func init() {
	env := utils.LoadEnv()

	region := env.AWSRegion
	profile := env.AWSProfile
	userPoolId := region + "_" + env.CognitoUserPoolId

	if region == "" || profile == "" {
		log.Fatalf("\033[33mcognito/config.go: \033[31mMissing AWS Environment Variables\033[0m\n\033[34mAWS_REGION := %s\nAWS_PROFILE := %s\033[0m\n", region, profile)
	}

	config = Config{
		Region:     region,
		Profile:    profile,
		UserPoolId: userPoolId,
	}
}

func GetConfig() Config {
	return config
}
