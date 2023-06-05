package api_home_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type HomeLambdaHandler struct {
	db *home_db_repo.Homes
}

// Create -
func Create(db *home_db_repo.Homes) *HomeLambdaHandler {
	return NewLambdaHandler(db)
}

// NewLambdaHandler -
func NewLambdaHandler(db *home_db_repo.Homes) *HomeLambdaHandler {
	return &HomeLambdaHandler{
		db: db,
	}
}
