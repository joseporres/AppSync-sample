package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


type Event struct{
	Id string `json:"id"`
	Sort string `json:"sort"`
}

type Output struct {
	Active bool `json:"active"`
	Icon string `json:"icon"`
	Title string `json:"title"`
	Url string `json:"url"`
}

func handler (ctx context.Context, event Event) (Output, error) {
	TABLE_NAME := os.Getenv("TABLENAME")


	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return Output{}, err
	}

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(event.Id),
			},
			"sort": {
				S: aws.String(event.Sort),
			},
		},
		
	})

	if err != nil {
		return Output{}, err
	}

	if result.Item == nil {
		return Output{Active:false, Icon:"", Title:"", Url:""}, nil
	}

	output := Output{}

	err1 := dynamodbattribute.UnmarshalMap(result.Item, &output)
	if err1 != nil {
		return Output{}, err1
	}

	return output, nil
}

func main (){
	lambda.Start(handler)
}