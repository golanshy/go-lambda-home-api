package config

import (
	"os"
)

type Config struct {
	helloMessage string
}

// NewConfigFromEnv -
func NewConfigFromEnv() *Config {

	return &Config{
		helloMessage: os.Getenv("HELLO_MESSAGE"),
	}
}
