package dtos

import "time"

type SensorDTO struct {
	API         string    `bson:"api,omitempty"`
	Id          string    `bson:"id,omitempty"`
	Name        string    `bson:"name,omitempty"`
	Description string    `bson:"description,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}
