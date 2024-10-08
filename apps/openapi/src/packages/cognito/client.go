package cognito

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var (
	client *cognitoidentityprovider.Client
)

func init() {
	cognitoConfig := GetConfig()

	config, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithRegion(cognitoConfig.Region),
		awsConfig.WithSharedConfigProfile(cognitoConfig.Profile),
	)

	if err != nil {
		panic(err)
	}

	client = cognitoidentityprovider.NewFromConfig(config)
}

func GetClient() *cognitoidentityprovider.Client {
	return client
}
