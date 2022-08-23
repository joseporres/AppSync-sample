package main

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type DynamoInt struct{
	N string `json:"N"`
}

type Result struct {
	AgePlus int `json:"agePlus"`
}

type Event struct{
	Age DynamoInt `json:"age"`
}

func handler(ctx context.Context, ev Event) (Result,error) {

	str := ev.Age.N
	sum, err := strconv.Atoi(str)
	if err != nil{
		return Result{},err
	}
	sum = sum +10
	res := Result{AgePlus: sum}
	return res, nil
}

func main (){
	lambda.Start(handler)
}