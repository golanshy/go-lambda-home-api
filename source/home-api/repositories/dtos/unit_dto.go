package dtos

import (
	"github.com/golanshy/go-lambda-home-api/data_models"
	"time"
)

type UnitDTO struct {
	API         string       `bson:"api,omitempty"`
	UnitId      string       `bson:"unit_id,omitempty"`
	Name        string       `bson:"name,omitempty"`
	Description string       `bson:"description,omitempty"`
	HomeId      string       `bson:"home_id,omitempty"`
	Sensors     []*SensorDTO `bson:"sensors,omitempty"`
	CreatedAt   time.Time    `bson:"created_at,omitempty"`
	UpdatedAt   time.Time    `bson:"updated_at,omitempty"`
}

func (d UnitDTO) ToUnit() *data_models.Unit {
	return &data_models.Unit{
		API:         d.API,
		UnitId:      d.UnitId,
		Name:        d.Name,
		Description: d.Description,
		HomeId:      d.HomeId,
		Sensors:     nil,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
