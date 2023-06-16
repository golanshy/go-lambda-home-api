package data_models

import "time"

type Unit struct {
	API         string    `json:"api,omitempty"`
	UnitId      string    `json:"unit_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	HomeId      string    `json:"home_id,omitempty"`
	Sensors     []*Sensor `json:"sensors,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
