package data_models

import "time"

type Temp struct {
	TempInC   float32   `json:"temp_in_c,omitempty"`
	TempInF   float32   `json:"temp_in_f,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type TempData struct {
	Data              []*Temp   `json:"data,omitempty"`
	StartTime         time.Time `json:"start_time,omitempty"`
	EndTime           time.Time `json:"end_time,omitempty"`
	ResolutionInHours int32     `json:"resolution_in_hours,omitempty"`
}
