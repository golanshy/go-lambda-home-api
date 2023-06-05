package config

import (
	"os"
)

type Config struct {
	HelloMessage string
	MongoURI     string
}

// NewConfigFromEnv -
func NewConfigFromEnv() *Config {

	return &Config{
		HelloMessage: os.Getenv("HELLO_MESSAGE"),
		MongoURI:     os.Getenv("MONGO_URI"),
	}
}
