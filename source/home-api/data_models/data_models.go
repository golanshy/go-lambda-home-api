package data_models

import "time"

type HomeData struct {
	API    string
	HomeId string
	Data   SensorsData
}

type SensorsData struct {
	Sensors []Sensor
}

type Sensor struct {
	SensorId string
	TempData []Temp
}

type Temp struct {
	Celsius    float32
	Fahrenheit float32
	CreatedAt  time.Time
}
