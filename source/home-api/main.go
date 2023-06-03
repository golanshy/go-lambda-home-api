package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golanshy/go-lambda-home-api/api_handler"
)

func main() {
	handler := api_handler.Create()
	lambda.Start(handler.HandleRequest)
}
