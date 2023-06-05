package api_hello_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type HelloLambdaHandler struct {
	dbClient *home_db_repo.StoreRepository
}

// Create -
func Create(dbClient *home_db_repo.StoreRepository) *HelloLambdaHandler {
	return NewLambdaHandler(dbClient)
}

// NewLambdaHandler -
func NewLambdaHandler(dbClient *home_db_repo.StoreRepository) *HelloLambdaHandler {
	return &HelloLambdaHandler{
		dbClient: dbClient,
	}
}
