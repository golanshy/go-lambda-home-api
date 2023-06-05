package api_sensor_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type SensorLambdaHandler struct {
	dbClient *home_db_repo.StoreRepository
}

// Create -
func Create(dbClient *home_db_repo.StoreRepository) *SensorLambdaHandler {
	return NewLambdaHandler(dbClient)
}

// NewLambdaHandler -
func NewLambdaHandler(dbClient *home_db_repo.StoreRepository) *SensorLambdaHandler {
	return &SensorLambdaHandler{
		dbClient: dbClient,
	}
}
