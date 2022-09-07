package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type App struct {
	Active bool   `json:"active"`
	Icon   string `json:"icon"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}

type Proccess struct {
	Active bool   `json:"active"`
	Icon   string `json:"icon"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}

type UsersId struct {
	UserId string `json:"id"`
}

func handler(ctx context.Context) (string, error) {

	TABLE_NAME := os.Getenv("TABLA_NAME")
	STAGE := os.Getenv("Stage")
	fmt.Println("Stage: ", STAGE)
	
	INDEX := "list-users"
	SORT := "SETTINGS"
	var usersId []UsersId

	//Lista de procesos, y lista de usuarios

	variableUrl := "https://webapp.ofvi." + STAGE +".protectasecuritycloud.pe/my_processes"
	
	activeArr := [8]bool{false, false, false, false, false, false, false, false}
	iconArr := [8]string{"Smartworking", "Metabase", "Client", "Platform", "Exa", "Onbase", "Software", "Comercial"}
	titleArr := [8]string{"Smart Working", "Metabase", "Cliente 360", "Plataforma Digital", "Exa", "Onbase", "Jira Software", "Sistema Comercial"}
	urlArr := [8]string{variableUrl, "https://metabase.core.prd.protectasecuritycloud.pe/", "https://servicios.protectasecurity.pe", "https://plataformadigital.protectasecurity.pe/ecommerce/extranet/login", "https://misecurity.exa.cl/login_standard/1", "https://workflow.protectasecurity.pe/AppNet/Login.aspx", "https://soporte.protectasecurity.pe/", "https://plataformadigital.protectasecurity.pe/sgc/"}

	var processArr []*dynamodb.AttributeValue

	process := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"active": {
				BOOL: aws.Bool(false),
			},
			"icon": {
				S: aws.String("Smartworking"),
			},
			"title": {
				S: aws.String("Smart Working"),
			},
			"url": {
				S: aws.String("/ofvi"),
			},
		},
	}

	processArr = append(processArr, process)


	var appArr []*dynamodb.AttributeValue
	for i := 0; i < 8; i++ {
		app := &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{
				"active": {
					BOOL: aws.Bool(activeArr[i]),
				},
				"icon": {
					S: aws.String(iconArr[i]),
				},
				"title": {
					S: aws.String(titleArr[i]),
				},
				"url": {
					S: aws.String(urlArr[i]),
				},
			},
		}
		appArr = append(appArr, app)
	}

	//Iniciar sesion en aws
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("us-east-1"))},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect session, %v", err))
	}
	svc := dynamodb.New(sess)

	errQ := svc.QueryPages(&dynamodb.QueryInput{
		TableName:              aws.String(TABLE_NAME),
		IndexName:              aws.String(INDEX),
		KeyConditionExpression: aws.String("sort = :sort"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":sort": {
				S: aws.String(SORT),
			},
		},
		ProjectionExpression: aws.String("id"),
	}, func(resultQuery *dynamodb.QueryOutput, last bool) bool {
		items := []UsersId{}
		err := dynamodbattribute.UnmarshalListOfMaps(resultQuery.Items, &items)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}

		usersId = append(usersId, items...)

		return true // keep paging
	})
	if errQ != nil {
		panic(fmt.Sprintf("Got error calling Query: %s", errQ))
	}

	for _, usr := range usersId {

		input := &dynamodb.UpdateItemInput{
			TableName: aws.String(TABLE_NAME),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(usr.UserId),
				},
				"sort": {
					S: aws.String(SORT),
				},
			},
			UpdateExpression: aws.String("set entryDate = :entryDate, creationDate = :creationDate, userStatus = :userStatus, userType = :userType, processes = :processes, apps = :apps"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":entryDate": {
					S: aws.String("2022-07-01"),
				},
				":creationDate": {
					S: aws.String("2022-07-01T00:00:01.001Z"),
				},
				":userStatus": {
					S: aws.String("ACTIVE"),
				},
				":userType": {
					S: aws.String("Usuario Protecta"),
				},
				":processes": {
					L: processArr,
				},
				":apps": {
					L: appArr,
				},	
			},
			ReturnValues: aws.String("UPDATED_NEW"),
		}
		_, err := svc.UpdateItem(input)
		if err != nil {
			panic(fmt.Sprintf("failed to Dynamodb Update Items, %v", err))
		}
	}

	

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":active": {
				BOOL: aws.Bool(false),
			},
			":icon": {
				S: aws.String("Smartworking"),
			},
			":title": {
				S: aws.String("Smart Working"),
			},
			":url": {
				S: aws.String("/ofvi"),
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
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET active = :active, icon = :icon, title = :title, #url = :url"),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		return "", err
	}


	// Agregar los objetos App y Process a la tabla
	for i := 0; i < 8; i++ {
		app := App{Active: activeArr[i], Icon: iconArr[i], Title: titleArr[i], Url: urlArr[i]}
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
					S: aws.String("APP-0" + strconv.Itoa(i+1)),
				},
			},
			ReturnValues:     aws.String("UPDATED_NEW"),
			UpdateExpression: aws.String("SET active = :active, icon = :icon, title = :title, #url = :url"),
		}

		_, err = svc.UpdateItem(input)

		if err != nil {
			return "", err
		}
	}

	return "SUCCESS", nil
}

func main() {
	lambda.Start(handler)
}