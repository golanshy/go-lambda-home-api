package api_sensor_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type SensorLambdaHandler struct {
	db *home_db_repo.Homes
}

// Create -
func Create(db *home_db_repo.Homes) *SensorLambdaHandler {
	return NewLambdaHandler(db)
}

// NewLambdaHandler -
func NewLambdaHandler(db *home_db_repo.Homes) *SensorLambdaHandler {
	return &SensorLambdaHandler{
		db: db,
	}
}
