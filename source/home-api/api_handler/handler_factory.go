package api_handler

import (
	"github.com/golanshy/go-lambda-home-api/api_hello_handler"
	"github.com/golanshy/go-lambda-home-api/api_home_handler"
	config2 "github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/handler"
)

// Create -
func Create() handler.Handler {
	config := config2.NewConfigFromEnv()
	helloHandler := api_hello_handler.Create(config)
	homeHandler := api_home_handler.Create(config)
	return NewLambdaHandler(config, helloHandler, homeHandler)
}
