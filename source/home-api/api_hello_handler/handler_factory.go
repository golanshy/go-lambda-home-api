package api_hello_handler

import (
	"github.com/golanshy/go-lambda-home-api/config"
)

type HelloLambdaHandler struct {
	config *config.Config
}

// Create -
func Create(c *config.Config) *HelloLambdaHandler {
	return NewLambdaHandler(c)
}

// NewLambdaHandler -
func NewLambdaHandler(c *config.Config) *HelloLambdaHandler {
	return &HelloLambdaHandler{
		config: c,
	}
}
