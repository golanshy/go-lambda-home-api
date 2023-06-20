package api_config_handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/go-lambda-home-api/handler"
	"net/http"
)

type ConfigResponse struct {
	EnableLocalLogs bool `json:"enable_local_logs"`
	DelayInSeconds  int  `json:"delay_in_seconds"`
}

func (l ConfigLambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

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

	configResponse := ConfigResponse{
		EnableLocalLogs: false,
		DelayInSeconds:  1 * 60 * 60, // 1 hour
	}
	response, err := json.Marshal(configResponse)

	res.StatusCode = http.StatusOK
	res.Body = string(response)

	return res, err
}
