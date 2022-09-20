package main

import (
	"context"
	"fmt"

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
	Username string `json:"username"`
	Password string `json:"password"`
}

func handler(ctx context.Context, event Event) (*cognito.AdminSetUserPasswordOutput, error) {

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
		appClientId:   "higbi45rjdihep9q4hujecbhg",
		userPoolId:    "us-east-1_Ruo06OeXI",
	}

	res,err :=client.cognitoClient.AdminSetUserPassword(&cognito.AdminSetUserPasswordInput{
		UserPoolId: aws.String(client.userPoolId),
		Username: aws.String(event.Username),
		Password: aws.String(event.Password),
		Permanent: aws.Bool(true),
	})
	
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
	
	return res,err
}

func main() {
	lambda.Start(handler)
}
