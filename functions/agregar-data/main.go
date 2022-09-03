package main

import (
	"context"
	"strconv"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type App struct {
	Active bool `json:"active"`
	Icon string `json:"icon"`
	Title string `json:"title"`
	Url string `json:"url"`
}


type Proccess struct {
	Active bool `json:"active"`
	Icon string `json:"icon"`
	Title string `json:"title"`
	Url string `json:"url"`
}
func handler(ctx context.Context)	 (string,error) {

	TABLE_NAME := os.Getenv("TABLENAME")
	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return "", err
	}

	svc := dynamodb.New(sess)

	proccess := Proccess{ Active: false, Icon:"" , Title: "Smart Working", Url: "/ofvi"}

	activeArr := [8] bool {false, false, false, false, false, false, false, false}
	iconArr := [8] string {"Client", "Metabase", "Client", "Platform", "Platform", "Onbase", "Software", "Comercial"}
	titleArr := [8] string {"Smart Working", "Metabase", "Cliente 360", "Plataforma Digital", "Exa", "Onbase", "Jira Software", "Sistema Comercial"}
	urlArr := [8] string {"https://webapp.ofvi.dev.protectasecuritycloud.pe/my_processes", "https://metabase.core.prd.protectasecuritycloud.pe/", "https://servicios.protectasecurity.pe", "https://plataformadigital.protectasecurity.pe/ecommerce/extranet/login", "https://misecurity.exa.cl/login_standard/1", "https://workflow.protectasecurity.pe/AppNet/Login.aspx", "https://soporte.protectasecurity.pe/", "https://plataformadigital.protectasecurity.pe/sgc/"}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":active": {
				BOOL: &proccess.Active,
			},
			":icon": {
				S: &proccess.Icon,
			},
			":title": {
				S: &proccess.Title,
			},
			":url": {
				S: &proccess.Url,
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#url": aws.String("url"),
		},
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("PROCESSES"),
			},
			"sort": {
				S: aws.String("PRC-01"),
			},
		},
		ReturnValues:    aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET active = :active, icon = :icon, title = :title, #url = :url"),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		return "", err
	}


	for i := 0; i < 8; i++ {
		app := App{ Active: activeArr[i], Icon:iconArr[i] , Title: titleArr[i], Url: urlArr[i]}
		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":active": {
					BOOL: &app.Active,
				},
				":icon": {
					S: &app.Icon,
				},
				":title": {
					S: &app.Title,
				},
				":url": {
					S: &app.Url,
				},
			},
			ExpressionAttributeNames: map[string]*string{
				"#url": aws.String("url"),
			},
			TableName: aws.String(TABLE_NAME),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String("APPLICATIONS"),
				},
				"sort": {
					S: aws.String("APP-0"+strconv.Itoa(i+1)),
				},
			},
			ReturnValues:    aws.String("UPDATED_NEW"),
			UpdateExpression: aws.String("SET active = :active, icon = :icon, title = :title, #url = :url"),
		}

		_, err = svc.UpdateItem(input)

		if err != nil {
			return "", err
		}
	}

	return "OK", nil
}

func main() {
	lambda.Start(handler)
}