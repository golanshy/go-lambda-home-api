package api_unit_handler

import (
	"github.com/golanshy/go-lambda-home-api/repositories/home_db_repo"
)

type UnitLambdaHandler struct {
	db *home_db_repo.Homes
}

// Create -
func Create(db *home_db_repo.Homes) *UnitLambdaHandler {
	return NewLambdaHandler(db)
}

// NewLambdaHandler -
func NewLambdaHandler(db *home_db_repo.Homes) *UnitLambdaHandler {
	return &UnitLambdaHandler{
		db: db,
	}
}
