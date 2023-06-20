package api_handler

import (
	"github.com/golanshy/go-lambda-home-api/api_config_handler"
	"github.com/golanshy/go-lambda-home-api/api_hello_handler"
	"github.com/golanshy/go-lambda-home-api/api_home_handler"
	"github.com/golanshy/go-lambda-home-api/api_sensor_handler"
	"github.com/golanshy/go-lambda-home-api/api_unit_handler"
	config "github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/handler"
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

// Create -
func Create() handler.Handler {
	cfg := config.NewConfigFromEnv()
	dbClient := home_db_repo.NewRepository(cfg)

	helloHandler := api_hello_handler.Create()
	configHandler := api_config_handler.Create()
	homeHandler := api_home_handler.Create(dbClient)
	unitHandler := api_unit_handler.Create(dbClient)
	sensorHandler := api_sensor_handler.Create(dbClient)

	return NewLambdaHandler(cfg, helloHandler, configHandler, homeHandler, unitHandler, sensorHandler)
}
