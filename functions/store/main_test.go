package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func TestHandlerReal(t *testing.T) {
	t.Run("success request", func(t *testing.T) {
		var svc dynamodbiface.DynamoDBAPI
		table := "otraTestTable2"
		os.Setenv("REGION", "us-east-1")
		e := `{
			"settings": {"name": "Jose", "docType": "DNI", "dni": "71112671", "gender": "Masculino", "birthDate": "18-09-01", "countryOfBirth": "Peru", "personalEmail": "jose.porres@devmente.com", "maritalStatus": "Soltero", "personalPhone": "999146656", "countryOfResidence": "Calle 2, Surco", "residenceDepartment": "Lima", "address": "calle 2 431", "area": "a", "subArea": "b", "workerType": "backend", "email": "jose.porresTESTING@devmente.com", "entryDate": "18-09-10", "phone": "123413241", "apps": [{"title": "a", "url": "bv", "icon": "v", "active": false}], "menu": [{"title": "a", "url": "b", "icon": "c", "active": false}], "processes": [{"title": "a", "url": "b", "icon": "c", "active": false}], "userType": "Usuario Externo", "role": "prueba", "days": 1, "homeOffice": 9, "photo": "link.com", "boss": "Jose", "bossName": "Jose", "user": "joseporres", "backup":"xxxx", "backupName":"listo"}
		  }`
		sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(os.Getenv("REGION"))}))	
		svc = dynamodb.New(sess)
		_, err := svc.CreateTable(&dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("sort"),
					AttributeType: aws.String("S"),
				},
			},
			BillingMode:            nil,
			GlobalSecondaryIndexes: nil,
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String(dynamodb.KeyTypeHash),
				},
				{
					AttributeName: aws.String("sort"),
					KeyType:       aws.String(dynamodb.KeyTypeRange),
				},
			},
			LocalSecondaryIndexes: nil,
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			SSESpecification:    nil,
			StreamSpecification: nil,
			TableName:           aws.String(table),
			Tags:                nil,
		})

		if err != nil {
			t.Fatal(err)
		}

		d := deps{
			ddb:   svc,
			table: os.Getenv("TableName"),
		}
		res := Event{}
		json.Unmarshal([]byte(e), &res)
		_, err1 := d.handler(context.TODO(), res)
		if err1 != nil {
			t.Fatal("Error")
		}

	})

}

