package main

import (
	"context"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)


type Option struct {
	Title string `json:"title"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
	Active bool `json:"active"`
}

type UserObject struct {
    Name string `json:"name"`
    DocType string `json:"docType"`
    Dni string `json:"dni"`
    Gender string `json:"gender"`
    BirthDate string `json:"birthDate"`
    CountryOfBirth string `json:"countryOfBirth"`
    PersonalEmail string `json:"personalEmail"`
    MaritalStatus string `json:"maritalStatus"`
    PersonalPhone string `json:"personalPhone"`
    CountryOfResidence string `json:"countryOfResidence"`
    ResidenceDepartment string `json:"residenceDepartment"`
    Address string `json:"address"`
    Area string `json:"area"`
    SubArea string `json:"subArea"`
    WorkerType string `json:"workerType"`
    Email string `json:"email"`
    EntryDate string `json:"entryDate"`
    Phone string `json:"phone"`
    Apps []Option `json:"apps"`
    Menu []Option `json:"menu"`
    Processes []Option `json:"processes"`
    UserType string `json:"userType"`
    Role string `json:"role"`
    Days int `json:"days"`
    HomeOffice int `json:"homeOffice"`
    Photo string `json:"photo"`
    Boss string `json:"boss"`
    BossName string `json:"bossName"`
    User string `json:"user"`
	Backup string `json:"backup"`
    BackupName string `json:"backupName"`
}

type UserInput struct {
	Id string `json:"id"`
	Sort string `json:"sort"`
    Name string `json:"name"`
    DocType string `json:"docType"`
    Dni string `json:"dni"`
    Gender string `json:"gender"`
    BirthDate string `json:"birthDate"`
    CountryOfBirth string `json:"countryOfBirth"`
    PersonalEmail string `json:"personalEmail"`
    MaritalStatus string `json:"maritalStatus"`
    PersonalPhone string `json:"personalPhone"`
    CountryOfResidence string `json:"countryOfResidence"`
    ResidenceDepartment string `json:"residenceDepartment"`
    Address string `json:"address"`
    Area string `json:"area"`
    SubArea string `json:"subArea"`
    WorkerType string `json:"workerType"`
    Email string `json:"email"`
    EntryDate string `json:"entryDate"`
    Phone string `json:"phone"`
    Apps []Option `json:"apps"`
    Menu []Option `json:"menu"`
    Processes []Option `json:"processes"`
    UserType string `json:"userType"`
    UserState string `json:"userState"`
    Role string `json:"role"`
    Days int `json:"days"`
    HomeOffice int `json:"homeOffice"`
    Photo string `json:"photo"`
    Boss string `json:"boss"`
    BossName string `json:"bossName"`
    User string `json:"user"`
	Backup string `json:"backup"`
    BackupName string `json:"backupName"`
	
}

type Event struct{
	Settings UserObject `json:"settings"`
}



func handler(ctx context.Context, event Event) (string, error) {
	TABLE_NAME := os.Getenv("SETTINGS_TABLE")

	var str string

	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return "", err
	}

	svc := dynamodb.New(sess)

	var appsArr []*dynamodb.AttributeValue
	var menuArr []*dynamodb.AttributeValue
	var processesArr []*dynamodb.AttributeValue

	for _, app := range event.Settings.Apps {
		attr := &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{
				"title": {
					S: aws.String(app.Title),
				},
				"url": {
					S: aws.String(app.Url),
				},
				"icon": {
					S: aws.String(app.Icon),
				},
				"active": {
					BOOL: aws.Bool(app.Active),
				},
			},
		}
		appsArr = append(appsArr, attr)
	}

	for _, menu := range event.Settings.Menu {
		attr := &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{
				"title": {
					S: aws.String(menu.Title),
				},
				"url": {
					S: aws.String(menu.Url),
				},
				"icon": {
					S: aws.String(menu.Icon),
				},
				"active": {
					BOOL: aws.Bool(menu.Active),
				},
			},
		}
		menuArr = append(menuArr, attr)
	}

	for _, process := range event.Settings.Processes {
		attr := &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{
				"title": {
					S: aws.String(process.Title),
				},
				"url": {
					S: aws.String(process.Url),
				},
				"icon": {
					S: aws.String(process.Icon),
				},
				"active": {
					BOOL: aws.Bool(process.Active),
				},
			},
		}
		processesArr = append(processesArr, attr)
	}



	input :=  &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": {
				S: aws.String(event.Settings.Name),
			},
			":docType": {
				S: aws.String(event.Settings.DocType),
			},
			":dni": {
				S: aws.String(event.Settings.Dni),
			},
			":gender": {
				S: aws.String(event.Settings.Gender),
			},
			":birthDate": {
				S: aws.String(event.Settings.BirthDate),
			},
			":countryOfBirth": {
				S: aws.String(event.Settings.CountryOfBirth),
			},
			":personalEmail": {
				S: aws.String(event.Settings.PersonalEmail),
			},
			":maritalStatus": {
				S: aws.String(event.Settings.MaritalStatus),
			},
			":personalPhone": {
				S: aws.String(event.Settings.PersonalPhone),
			},
			":countryOfResidence": {
				S: aws.String(event.Settings.CountryOfResidence),
			},
			":residenceDepartment": {
				S: aws.String(event.Settings.ResidenceDepartment),
			},
			":address": {
				S: aws.String(event.Settings.Address),
			},
			":area": {
				S: aws.String(event.Settings.Area),
			},
			":subArea": {
				S: aws.String(event.Settings.SubArea),
			},
			":workerType": {
				S: aws.String(event.Settings.WorkerType),
			},
			":email": {
				S: aws.String(event.Settings.Email),
			},
			":entryDate": {
				S: aws.String(event.Settings.EntryDate),
			},
			":phone": {
				S: aws.String(event.Settings.Phone),
			},
			":apps": {
				L: appsArr,
			},
			":menu": {
				L: menuArr,
			},
			":processes": {
				L: processesArr,
			},
			":userType": {
				S: aws.String(event.Settings.UserType),
			},
			":userState": {
				S: aws.String("UNCONFIRMED"),
			},
			":role": {
				S: aws.String(event.Settings.Role),
			},
			":days": {
				N: aws.String(strconv.Itoa(event.Settings.Days)),
			},
			":homeOffice": {
				N: aws.String(strconv.Itoa(event.Settings.HomeOffice)),
			},
			":photo": {
				S: aws.String(event.Settings.Photo),
			},
			":boss": {
				S: aws.String(event.Settings.Boss),
			},
			":bossName": {
				S: aws.String(event.Settings.BossName),
			},
			":user": {
				S: aws.String(event.Settings.User),
			},
			":backup": {
				S: aws.String(event.Settings.Backup),
			},
			":backupName": {
				S: aws.String(event.Settings.BackupName),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("name"),
			"#role": aws.String("role"),
			"#user": aws.String("user"),
			"#backup": aws.String("backup"),
		},
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(event.Settings.Email),
			},
			"sort": {
				S: aws.String("SETTINGS"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET #name = :name, docType = :docType, dni = :dni, gender = :gender, birthDate = :birthDate, countryOfBirth = :countryOfBirth, personalEmail = :personalEmail, maritalStatus = :maritalStatus, personalPhone = :personalPhone, countryOfResidence = :countryOfResidence, residenceDepartment = :residenceDepartment, address = :address, area = :area, subArea = :subArea, workerType = :workerType, email = :email, entryDate = :entryDate, phone = :phone, apps = :apps, menu = :menu, processes = :processes, userType = :userType, userState = :userState, #role = :role, days = :days, homeOffice = :homeOffice, photo = :photo, boss = :boss, bossName = :bossName, #user = :user, #backup = :backup, backupName = :backupName"),
	}
	_, err = svc.UpdateItem(input)

	if err != nil {
		return "", err
	}

	str = "Success"
	return str, nil
}



func main () {
	lambda.Start(handler)
}