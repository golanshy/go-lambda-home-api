package data_models

import "time"

type TimeSeries struct {
	TimeDateStamp  []time.Time      `json:"time_date_stamp,omitempty"`
	TimeStamp      []string         `json:"time_stamp,omitempty"`
	TimeSeriesData []TimeSeriesData `json:"time_series_data,omitempty"`
}

type TimeSeriesData struct {
	SensorId                 string    `json:"sensor_id"`
	SensorName               string    `json:"sensor_name"`
	SensorDescription        string    `json:"sensor_description"`
	TempReadingsInC          []float32 `json:"temp_readings_in_c,omitempty"`
	TempReadingsInPercentage []float32 `json:"temp_readings_in_percentage,omitempty"`
}
