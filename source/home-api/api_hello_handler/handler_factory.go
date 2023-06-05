package api_hello_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type HelloLambdaHandler struct {
	db *home_db_repo.Homes
}

// Create -
func Create(db *home_db_repo.Homes) *HelloLambdaHandler {
	return NewLambdaHandler(db)
}

// NewLambdaHandler -
func NewLambdaHandler(db *home_db_repo.Homes) *HelloLambdaHandler {
	return &HelloLambdaHandler{
		db: db,
	}
}
