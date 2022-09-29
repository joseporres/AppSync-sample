package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type awsCognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
	userPoolId    string
}

type Event struct {
	Email string `json:"email"`
	
}

type Response struct {
	Username string `json:"username"`
	Enabled bool `json:"enabled"`
	AccountStatus string `json:"accountStatus"`
	Email string `json:"email"`
	EmailVerified string `json:"emailVerified"`
	PhoneNumberVerified string `json:"phoneNumberVerified"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}


func handler(ctx context.Context, event Event) ([]Response, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(*aws.String("us-east-1"))},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect session, %v", err))
	}
	// INICIAR SESSION EN COGNITO
	svc := cognito.New(sess)

	client := awsCognitoClient{
		cognitoClient: svc,
		appClientId:   "4m0fh965du3rhl0cvgtova3p3l",
		userPoolId:    "us-east-1_uBHINArtY",
	}

	res, err := client.cognitoClient.ListUsers(&cognito.ListUsersInput{
		Filter: aws.String("email = \"" + event.Email + "\""),
		UserPoolId: aws.String(client.userPoolId),
	})

	if err != nil {
        fmt.Println("Got error listing users")
        os.Exit(1)
    }

	var response []Response
	for _, user := range res.Users {
		fmt.Println("user: ", user)
		response = append(response, Response{
			Username: *user.Username,
			Enabled: *user.Enabled,
			AccountStatus: *user.UserStatus,
			Email: *user.Attributes[0].Value,
			Updated: *user.UserLastModifiedDate,
			Created: *user.UserCreateDate,
		})
	}

	return response, nil



}

func main() {
	lambda.Start(handler)
}