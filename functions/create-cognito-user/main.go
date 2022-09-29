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

type Response struct {
	Username            string    `json:"username"`
	Enabled             bool      `json:"enabled"`
	AccountStatus       string    `json:"accountStatus"`
	Email               string    `json:"email"`
	EmailVerified       string    `json:"emailVerified"`
	PhoneNumberVerified string    `json:"phoneNumberVerified"`
	Updated             time.Time `json:"updated"`
	Created             time.Time `json:"created"`
}

type Event struct {
	Email            string `json:"email"`
	Name             string `json:"name"`
}

func handler(ctx context.Context, event Event)(string,error){
	var result string
	var response []Response
	// CONECTAR SESSION CON AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(*aws.String("us-east-1"))},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect session, %v", err))
	}
	// INICIAR SESSION EN COGNITO
	svc := cognito.New(sess)

	fmt.Println("APP_CLIENT_ID : ", os.Getenv("APP_CLIENT_ID"))
	fmt.Println("USER_POOL_ID : ", os.Getenv("USER_POOL_ID"))

	client := awsCognitoClient{
		cognitoClient: svc,
		appClientId:   os.Getenv("APP_CLIENT_ID"),
		userPoolId:    os.Getenv("USER_POOL_ID"),
	}

	result, err = client.AdminCreateUser(event.Email, event.Name)

	if err != nil {
		fmt.Println("Error :", err)
		return "", err
	}
	fmt.Println("CLIENTE :", client)
	fmt.Println("Response :", response)
	fmt.Println("result :", result)

	return result, nil


}

func (ctx *awsCognitoClient) AdminCreateUser(email string, name string) (string, error) {

	user := &cognito.AdminCreateUserInput{
		UserPoolId: aws.String(ctx.userPoolId),
		Username:   aws.String(email),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(name),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
	}
	fmt.Println("USER: ", user)

	result, err := ctx.cognitoClient.AdminCreateUser(user)
	if err != nil {
		fmt.Println("Error : AdminCreateUser", err)
		return "", err
	}
	return *result.User.Username, nil
}

func main() {
	lambda.Start(handler)
}