package handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

// Handler - interface
type Handler interface {
	Run(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (Response, error)
}

type lambdaHandler struct {
	helloeMessage string
}

type LambdaResponse struct {
	Message string
}

func (l lambdaHandler) Run(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (Response, error) {
	lambdaResponse := LambdaResponse{
		Message: l.helloeMessage,
	}

	response, err := json.Marshal(lambdaResponse)

	res := Response{
		StatusCode:      200,
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

// NewLambdaHandler -
func NewLambdaHandler(
	helloeMessage string,
) *lambdaHandler {
	return &lambdaHandler{
		helloeMessage: helloeMessage,
	}
}
