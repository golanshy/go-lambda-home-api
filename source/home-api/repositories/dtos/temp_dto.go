package dtos

import "time"

type TempRecordDTO struct {
	HomeId      string    `bson:"home_id,omitempty"`
	UnitId      string    `bson:"unit_id,omitempty"`
	SensorId    string    `bson:"sensor_id,omitempty"`
	Temperature float32   `bson:"temperature,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}

type TempDataDTO struct {
	Data              []TempRecordDTO `bson:"data,omitempty"`
	StartTime         time.Time       `bson:"start_time,omitempty"`
	EndTime           time.Time       `bson:"end_time,omitempty"`
	ResolutionInHours int32           `bson:"resolution_in_hours,omitempty"`
}
