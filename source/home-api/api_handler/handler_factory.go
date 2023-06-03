package api_handler

import (
	config2 "github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/handler"
	"github.com/golanshy/go-lambda-home-api/hello_handler"
	"github.com/golanshy/go-lambda-home-api/home_handler"
)

// Create -
func Create() handler.Handler {
	config := config2.NewConfigFromEnv()
	helloHandler := hello_handler.Create(config)
	homeHandler := home_handler.Create(config)
	return NewLambdaHandler(config, helloHandler, homeHandler)
}
