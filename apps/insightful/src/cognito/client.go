package cognito

import (
	"context"
	"log"
	"sync"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var (
	client *cognitoidentityprovider.Client
)
var clientDebug = "\033[33mcognito/client.go\033[0m"
var onceInitClient sync.Once

func initClient() {
	cognitoConfig := GetConfig()

	config, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithLogConfigurationWarnings(true),
		awsConfig.WithRegion(cognitoConfig.Region),
		awsConfig.WithSharedConfigProfile(cognitoConfig.Profile),
	)

	if err != nil {
		log.Printf("%s CognitoConfig: %#v\n", clientDebug, cognitoConfig)
		log.Fatalf("%s Failed to create config: %v", clientDebug, err)
	}

	client = cognitoidentityprovider.NewFromConfig(config)
}

func GetClient() *cognitoidentityprovider.Client {
	onceInitClient.Do(initClient)

	return client
}
