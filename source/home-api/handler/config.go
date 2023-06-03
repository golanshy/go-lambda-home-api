package handler

import (
	"os"
)

type config struct {
	helloMessage string
}

// NewConfigFromEnv -
func NewConfigFromEnv() *config {

	return &config{
		helloMessage: os.Getenv("HELLO_MESSAGE"),
	}
}
