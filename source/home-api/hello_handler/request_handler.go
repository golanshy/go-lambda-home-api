package hello_handler

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

func (l HelloLambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

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

	name, ok := req.QueryStringParameters["name"]

	if !ok {
		lambdaResponse := LambdaResponse{
			Message: "unknown",
		}
		response, err := json.Marshal(lambdaResponse)

		res.StatusCode = http.StatusBadRequest
		res.Body = string(response)
		return res, err
	}

	lambdaResponse := LambdaResponse{
		Message: fmt.Sprintf("Hello %s!", name),
	}
	response, err := json.Marshal(lambdaResponse)

	res.StatusCode = http.StatusOK
	res.Body = string(response)

	return res, err
}
