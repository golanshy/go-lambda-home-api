package dtos

import (
	"github.com/golanshy/go-lambda-home-api/data_models"
	"time"
)

type TempRecordDTO struct {
	HomeId      string    `bson:"home_id,omitempty"`
	UnitId      string    `bson:"unit_id,omitempty"`
	SensorId    string    `bson:"sensor_id,omitempty"`
	Temperature float32   `bson:"temperature,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}

type TempDataDTO struct {
	Data              []TempRecordDTO `bson:"data,omitempty"`
	StartTime         time.Time       `bson:"start_time,omitempty"`
	EndTime           time.Time       `bson:"end_time,omitempty"`
	ResolutionInHours int32           `bson:"resolution_in_hours,omitempty"`
}

func FromTempDTOs(values []*TempRecordDTO) []*data_models.Temp {
	if values == nil {
		return nil
	}
	result := make([]*data_models.Temp, 0)
	for _, value := range values {
		if value != nil {
			result = append(result, fromTempDTO(value))
		}
	}
	return result
}

func fromTempDTO(value *TempRecordDTO) *data_models.Temp {
	return &data_models.Temp{
		TempInC:   value.Temperature,
		TempInF:   (value.Temperature * 1.8) + 32.0,
		CreatedAt: value.CreatedAt,
	}
}
