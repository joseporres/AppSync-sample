package main

import (
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Object struct {
	Id string `json:"id"`
	Sort string `json:"sort"`
	Nombre string `json:"nombre"`
}

func handler(event Object) (string, error) {
	TABLE_NAME := os.Getenv("TABLENAME")

	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return "", err
	}

	svc := dynamodb.New(sess)
	object := Object{"Put", event.Sort, event.Nombre}
	item, err := dynamodbattribute.MarshalMap(object)
	if err != nil {
		return "", err
	}
	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(TABLE_NAME),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		return "", err
	}


	input2 := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":nombre": {
				S: aws.String(event.Nombre),
			},
		},
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("Update"),
			},
			"sort": {
				S: aws.String(event.Sort),
			},
		},
		ReturnValues:    aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET nombre = :nombre"),

	
	}

	_, err = svc.UpdateItem(input2)
	if err != nil {
		return "", err
	}
	
	return "Item saved with put and update", nil
}




func main () {
	lambda.Start(handler)
	
}