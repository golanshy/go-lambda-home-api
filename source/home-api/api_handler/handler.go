package api_handler

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	config2 "github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/handler"
	"strings"
)

var (
	homePath  = "home"
	helloPath = "hello"
	notFound  = "not_found"
)

type LambdaHandler struct {
	config       *config2.Config
	helloHandler handler.Handler
	homeHandler  handler.Handler
}

type LambdaResponse struct {
	Message string
}

// NewLambdaHandler -
func NewLambdaHandler(c *config2.Config, helloHandler handler.Handler, homeHandler handler.Handler) *LambdaHandler {
	return &LambdaHandler{
		config:       c,
		helloHandler: helloHandler,
		homeHandler:  homeHandler,
	}
}

func (l LambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

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

	if strings.Contains(req.Path, helloPath) {
		return l.helloHandler.HandleRequest(ctx, req)
	} else if strings.Contains(req.Path, homePath) {
		return l.homeHandler.HandleRequest(ctx, req)
	} else {
		return res, errors.New(notFound)
	}
}
