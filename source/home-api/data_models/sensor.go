package data_models

import "time"

type Sensor struct {
	API         string    `json:"api,omitempty"`
	SensorId    string    `json:"sensor_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	TempData    *TempData `json:"temp_data,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
