package api_handler

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/go-lambda-home-api/api_sensor_handler"
	"github.com/golanshy/go-lambda-home-api/api_unit_handler"
	config2 "github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/handler"
	"log"
	"strings"
)

var (
	helloPath  = "hello"
	configPath = "config"
	homePath   = "home"
	unitPath   = "unit"
	sensorPath = "sensor"
	notFound   = "not_found"
)

type LambdaHandler struct {
	config        *config2.Config
	helloHandler  handler.Handler
	configHandler handler.Handler
	homeHandler   handler.Handler
	unitHandler   handler.Handler
	sensorHandler handler.Handler
}

type LambdaResponse struct {
	Message string
}

// NewLambdaHandler -
func NewLambdaHandler(c *config2.Config, helloHandler handler.Handler, configHandler handler.Handler, homeHandler handler.Handler, unitHandler *api_unit_handler.UnitLambdaHandler, sensorHandler *api_sensor_handler.SensorLambdaHandler) *LambdaHandler {
	return &LambdaHandler{
		config:        c,
		helloHandler:  helloHandler,
		configHandler: configHandler,
		homeHandler:   homeHandler,
		unitHandler:   unitHandler,
		sensorHandler: sensorHandler,
	}
}

func (l LambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

	log.Printf("HandleRequest Path: %s", req.Path)

	if strings.Contains(req.Path, helloPath) {
		return l.helloHandler.HandleRequest(ctx, req)
	} else if strings.Contains(req.Path, configPath) {
		return l.configHandler.HandleRequest(ctx, req)
	} else if strings.Contains(req.Path, homePath) {
		return l.homeHandler.HandleRequest(ctx, req)
	} else if strings.Contains(req.Path, unitPath) {
		return l.unitHandler.HandleRequest(ctx, req)
	} else if strings.Contains(req.Path, sensorPath) {
		return l.sensorHandler.HandleRequest(ctx, req)
	} else {

		res := handler.Response{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
				"Cache-Control":                    "no-cache; no-store",
				"Content-Type":                     "application/json",
				"Accept":                           "application/json",
				"Content-Security-Policy":          "default-src self",
				"Strict-Transport-Security":        "max-age=31536000; includeSubDomains",
				"X-Content-Type-Options":           "nosniff",
				"X-XSS-Protection":                 "1; mode=block",
				"X-Frame-Options":                  "DENY",
			},
		}
		return res, errors.New(notFound)
	}
}
