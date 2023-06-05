package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golanshy/go-lambda-home-api/api_handler"
	"log"
)

func main() {
	log.Printf("application started")
	handler := api_handler.Create()
	lambda.Start(handler.HandleRequest)
}
