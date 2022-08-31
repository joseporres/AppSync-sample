package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
    LastSessionDate string `json:"lastSessionDate"`
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
    LastSessionDate string `json:"lastSessionDate"`
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
	in := UserInput{ 
		Id: event.Settings.Email,
		Sort: "SETTINGS",
		Name: event.Settings.Name,
		DocType: event.Settings.DocType,
		Dni: event.Settings.Dni,
		Gender: event.Settings.Gender,
		BirthDate: event.Settings.BirthDate,
		CountryOfBirth: event.Settings.CountryOfBirth,
		PersonalEmail: event.Settings.PersonalEmail,
		MaritalStatus: event.Settings.MaritalStatus,
		PersonalPhone: event.Settings.PersonalPhone,
		CountryOfResidence: event.Settings.CountryOfResidence,
		ResidenceDepartment: event.Settings.ResidenceDepartment,
		Address: event.Settings.Address,
		Area: event.Settings.Area,
		SubArea: event.Settings.SubArea,
		WorkerType: event.Settings.WorkerType,
		Email: event.Settings.Email,
		EntryDate: event.Settings.EntryDate,
		LastSessionDate: event.Settings.LastSessionDate,
		Phone: event.Settings.Phone,
		Apps: event.Settings.Apps,
		Menu: event.Settings.Menu,
		Processes: event.Settings.Processes,
		UserType: event.Settings.UserType,
		UserState: event.Settings.UserState,
		Role: event.Settings.Role,
		Days: event.Settings.Days,
		HomeOffice: event.Settings.HomeOffice,
		Photo: event.Settings.Photo,
		Boss: event.Settings.Boss,
		BossName: event.Settings.BossName,
		User: event.Settings.User,
	}

	fmt.Println("Evento: ", event)


	fmt.Println("In: ", in)

	item, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return "", err
	}

	fmt.Println("Marshall: ",item)

	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(TABLE_NAME),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return "", err
	}

	str = "Success"
	return str, nil
}



func main () {
	lambda.Start(handler)
}