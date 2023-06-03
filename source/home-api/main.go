package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golanshy/go-lambda-home-api/source/hello/handler"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}
