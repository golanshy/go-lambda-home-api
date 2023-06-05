package api_unit_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type UnitLambdaHandler struct {
	dbClient *home_db_repo.StoreRepository
}

// Create -
func Create(dbClient *home_db_repo.StoreRepository) *UnitLambdaHandler {
	return NewLambdaHandler(dbClient)
}

// NewLambdaHandler -
func NewLambdaHandler(dbClient *home_db_repo.StoreRepository) *UnitLambdaHandler {
	return &UnitLambdaHandler{
		dbClient: dbClient,
	}
}
