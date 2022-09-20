package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


type UserAttributesType struct {
	CognitoUser       string `json:"cognito:user_status"`
	Nombre            string `json:"custom:nombre"`
	Email             string `json:"email"`
	Sub               string `json:"sub"`
	UsernameParameter string `json:"usernameParameter"`
}

type CognitoEventUserPoolsCallerContext struct {
	AWSSDKVersion string `json:"awsSdkVersion"`
	ClientID      string `json:"clientId"`
}

type CognitoEventUserPoolsHeader struct {
	Version       string                             `json:"version"`
	TriggerSource string                             `json:"triggerSource"`
	Region        string                             `json:"region"`
	UserPoolID    string                             `json:"userPoolId"`
	CallerContext CognitoEventUserPoolsCallerContext `json:"callerContext"`
	UserName      string                             `json:"userName"`
}

// CognitoEventUserPoolsCustomMessage is sent by AWS Cognito User Pools before a verification or MFA message is sent,
// allowing a user to customize the message dynamically.
type CognitoEventUserPoolsCustomMessage struct {
	CognitoEventUserPoolsHeader
	Request  CognitoEventUserPoolsCustomMessageRequest  `json:"request"`
	Response CognitoEventUserPoolsCustomMessageResponse `json:"response"`
}

// CognitoEventUserPoolsCustomMessageRequest contains the request portion of a CustomMessage event
type CognitoEventUserPoolsCustomMessageRequest struct {
	UserAttributes    UserAttributesType `json:"userAttributes"`
	CodeParameter     string             `json:"codeParameter"`
	UsernameParameter string             `json:"usernameParameter"`
	ClientMetadata    map[string]string  `json:"clientMetadata"`
}

// CognitoEventUserPoolsCustomMessageResponse contains the response portion of a CustomMessage event
type CognitoEventUserPoolsCustomMessageResponse struct {
	SMSMessage   string `json:"smsMessage"`
	EmailMessage string `json:"emailMessage"`
	EmailSubject string `json:"emailSubject"`
}

func handler(ctx context.Context, event CognitoEventUserPoolsCustomMessage) (CognitoEventUserPoolsCustomMessage, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	if err != nil {
		fmt.Println(err.Error())
		return CognitoEventUserPoolsCustomMessage{}, err
	}
	svc := s3.New(sess)
	
	var htmlPage string

	htmlPage = "externo.html"
	
	rawObject, err := svc.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("BucketName")),
			Key:    aws.String(htmlPage),
		})
	if err != nil {
		fmt.Println(err.Error())
		return CognitoEventUserPoolsCustomMessage{}, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(rawObject.Body)
	result := buf.String()

	t, err := template.New("mailhtml").Parse(result)
	if err != nil {
		fmt.Println(err.Error())
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, event.Request.UserAttributes); err != nil {
		fmt.Println(err.Error())
	}

	resultT := tpl.String()
	event.Response.EmailMessage = resultT
	event.Response.EmailSubject = "Invitación a Oficina Virtual: Cuenta Nueva"
	event.Response.SMSMessage = ""

	return event, nil

}

func main() {
	lambda.Start(handler)
}
