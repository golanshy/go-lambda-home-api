package api_home_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/go-lambda-home-api/handler"
	"net/http"
)

type LambdaResponse struct {
	Message string
}

func (l HomeLambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

	res := handler.Response{
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

	switch req.HTTPMethod {
	case http.MethodPost:
		return postHome(req, res)
	case http.MethodGet:
		return getHome(req, res)
	}

	lambdaResponse := LambdaResponse{
		Message: "not_found",
	}
	response, err := json.Marshal(lambdaResponse)
	res.StatusCode = http.StatusNotFound
	res.Body = string(response)
	return res, err
}

func getHome(req events.APIGatewayProxyRequest, res handler.Response) (handler.Response, error) {
	name, ok := req.QueryStringParameters["name"]

	if !ok {
		lambdaResponse := LambdaResponse{
			Message: "Welcome home!",
		}
		response, err := json.Marshal(lambdaResponse)

		res.StatusCode = http.StatusOK
		res.Body = string(response)
		return res, err
	}

	lambdaResponse := LambdaResponse{
		Message: fmt.Sprintf("Welcome home, %s!", name),
	}
	response, err := json.Marshal(lambdaResponse)
	res.StatusCode = http.StatusOK
	res.Body = string(response)
	return res, err
}

func postHome(req events.APIGatewayProxyRequest, res handler.Response) (handler.Response, error) {
	lambdaResponse := LambdaResponse{
		Message: "Post home!",
	}
	response, err := json.Marshal(lambdaResponse)

	res.StatusCode = http.StatusOK
	res.Body = string(response)
	return res, err
}
