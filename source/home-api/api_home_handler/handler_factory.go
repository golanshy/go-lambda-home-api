package api_home_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type HomeLambdaHandler struct {
	dbClient *home_db_repo.StoreRepository
}

// Create -
func Create(dbClient *home_db_repo.StoreRepository) *HomeLambdaHandler {
	return NewLambdaHandler(dbClient)
}

// NewLambdaHandler -
func NewLambdaHandler(dbClient *home_db_repo.StoreRepository) *HomeLambdaHandler {
	return &HomeLambdaHandler{
		dbClient: dbClient,
	}
}
