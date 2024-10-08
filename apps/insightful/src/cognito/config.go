package cognito

import (
	"fmt"
	"insightful/src/utils"
	"log"
)

type Config struct {
	Region     string
	Profile    string
	UserPoolId string
}

var config Config
var configDebug = "\033[33mcognito/config.go\033[0m"

var UserPoolId = config.UserPoolId

func init() {
	env := utils.LoadEnv()

	region := env.AWSRegion
	profile := env.AWSProfile
	cognitoUserPoolId := env.CognitoUserPoolId
	userPoolId := region + "_" + cognitoUserPoolId

	if region == "" || profile == "" || cognitoUserPoolId == "" {
		regionVar := fmt.Sprintf("AWSRegion = %s", region)
		profileVar := fmt.Sprintf("AWSProfile = %s", profile)
		userPoolVar := fmt.Sprintf("CognitoUserPoolId = %s", cognitoUserPoolId)

		log.Fatalf("%s Missing Cognito Configuration\n\033[34m%s\n%s\n%s\n\033[0m", configDebug, regionVar, profileVar, userPoolVar)
	}

	config = Config{
		Region:     region,
		Profile:    profile,
		UserPoolId: userPoolId,
	}
}

func GetConfig() *Config {
	return &config
}
