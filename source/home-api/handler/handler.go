package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response events.APIGatewayProxyResponse

// Handler - interface
type Handler interface {
	Run(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (Response, error)
	HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error)
}

type LambdaHandler struct {
	helloMessage string
}

type LambdaResponse struct {
	Message string
}

func (l LambdaHandler) Run(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (Response, error) {

	lambdaResponse := LambdaResponse{
		Message: l.helloMessage,
	}

	response, err := json.Marshal(lambdaResponse)

	res := Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Cache-Control":                    "no-cache; no-store",
			"Content-Type":                     "application/json",
			"Content-Security-Policy":          "default-src self",
			"Strict-Transport-Security":        "max-age=31536000; includeSubDomains",
			"X-Content-Type-Options":           "nosniff",
			"X-XSS-Protection":                 "1; mode=block",
			"X-Frame-Options":                  "DENY",
		},
		Body: string(response),
	}

	return res, err
}

func (l LambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {

	name, ok := req.QueryStringParameters["name"]

	res := Response{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Cache-Control":                    "no-cache; no-store",
			"Content-Type":                     "application/json",
			"Content-Security-Policy":          "default-src self",
			"Strict-Transport-Security":        "max-age=31536000; includeSubDomains",
			"X-Content-Type-Options":           "nosniff",
			"X-XSS-Protection":                 "1; mode=block",
			"X-Frame-Options":                  "DENY",
		},
	}

	if !ok {
		lambdaResponse := LambdaResponse{
			Message: "Unknown",
		}
		response, err := json.Marshal(lambdaResponse)

		res.StatusCode = http.StatusBadRequest
		res.Body = string(response)
		return res, err
	}

	lambdaResponse := LambdaResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}
	response, err := json.Marshal(lambdaResponse)

	res.StatusCode = http.StatusOK
	res.Body = string(response)

	return res, err
}

// NewLambdaHandler -
func NewLambdaHandler(
	helloMessage string,
) *LambdaHandler {
	return &LambdaHandler{
		helloMessage: helloMessage,
	}
}
