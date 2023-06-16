package dtos

import (
	"github.com/golanshy/go-lambda-home-api/data_models"
	"time"
)

type SensorDTO struct {
	API         string    `bson:"api,omitempty"`
	SensorId    string    `bson:"sensor_id,omitempty"`
	Name        string    `bson:"name,omitempty"`
	Description string    `bson:"description,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}

func FromSensors(sensors []*data_models.Sensor) []*SensorDTO {
	if sensors == nil {
		return nil
	}
	result := make([]*SensorDTO, 0)
	for _, sensor := range sensors {
		if sensor != nil {
			result = append(result, fromSensor(sensor))
		}
	}
	return result
}

func fromSensor(sensor *data_models.Sensor) *SensorDTO {
	return &SensorDTO{
		API:         sensor.API,
		SensorId:    sensor.SensorId,
		Name:        sensor.Name,
		Description: sensor.Description,
		CreatedAt:   sensor.CreatedAt,
		UpdatedAt:   sensor.UpdatedAt,
	}
}

func FromSensorDTOs(sensors []*SensorDTO) []*data_models.Sensor {
	if sensors == nil {
		return nil
	}
	result := make([]*data_models.Sensor, 0)
	for _, sensor := range sensors {
		if sensor != nil {
			result = append(result, fromSensorDTO(sensor))
		}
	}
	return result
}

func fromSensorDTO(sensor *SensorDTO) *data_models.Sensor {
	return &data_models.Sensor{
		API:         sensor.API,
		SensorId:    sensor.SensorId,
		Name:        sensor.Name,
		Description: sensor.Description,
		CreatedAt:   sensor.CreatedAt,
		UpdatedAt:   sensor.UpdatedAt,
	}
}
