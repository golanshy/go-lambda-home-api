package config

import (
	"os"
)

type Config struct {
	HelloMessage string
}

// NewConfigFromEnv -
func NewConfigFromEnv() *Config {

	return &Config{
		HelloMessage: os.Getenv("HELLO_MESSAGE"),
	}
}
