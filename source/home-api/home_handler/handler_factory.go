package home_handler

import (
	"github.com/golanshy/go-lambda-home-api/config"
)

type HomeLambdaHandler struct {
	config *config.Config
}

// Create -
func Create(c *config.Config) *HomeLambdaHandler {
	return NewLambdaHandler(c)
}

// NewLambdaHandler -
func NewLambdaHandler(c *config.Config) *HomeLambdaHandler {
	return &HomeLambdaHandler{
		config: c,
	}
}
