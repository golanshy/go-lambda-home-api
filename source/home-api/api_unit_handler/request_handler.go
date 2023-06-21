package api_unit_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/go-lambda-home-api/data_models"
	"github.com/golanshy/go-lambda-home-api/handler"
	"log"
	"math"
	"net/http"
	"sort"
	"time"
)

type LambdaResponse struct {
	Message string
}

func (l UnitLambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

	res := handler.Response{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Cache-Control":                    "no-cache; no-store",
			"Accept":                           "application/json",
			"Content-Type":                     "application/json",
			"Content-Security-Policy":          "default-src self",
			"Strict-Transport-Security":        "max-age=31536000; includeSubDomains",
			"X-Content-Type-Options":           "nosniff",
			"X-XSS-Protection":                 "1; mode=block",
			"X-Frame-Options":                  "DENY",
		},
	}

	switch req.HTTPMethod {
	case http.MethodGet:
		return l.getUnit(ctx, req, res)
	case http.MethodPost:
		return l.insertUnitData(ctx, req, res)
	case http.MethodPut:
		return l.updateUnit(ctx, req, res)
	}

	lambdaResponse := LambdaResponse{
		Message: "not_found",
	}
	response, err := json.Marshal(lambdaResponse)
	res.StatusCode = http.StatusNotFound
	res.Body = string(response)
	return res, err
}

func (l UnitLambdaHandler) getUnit(ctx context.Context, req events.APIGatewayProxyRequest, res handler.Response) (handler.Response, error) {
	id, ok := req.QueryStringParameters["id"]

	if !ok {
		lambdaResponse := LambdaResponse{
			Message: "unit id missing",
		}
		response, err := json.Marshal(lambdaResponse)

		res.StatusCode = http.StatusBadRequest
		res.Body = string(response)
		return res, err
	}

	unitDate, err := l.dbClient.GetUnitData(ctx, id)
	if err != nil {
		log.Printf("InsertUnitData error: %s", err.Error())
		lambdaResponse := LambdaResponse{
			Message: fmt.Sprintf("failed reading unit data"),
		}
		response, _ := json.Marshal(lambdaResponse)
		res.StatusCode = http.StatusInternalServerError
		res.Body = string(response)
		return res, err
	}

	sort.Slice(unitDate.Sensors, func(i, j int) bool {
		return len(unitDate.Sensors[i].SensorId) <= len(unitDate.Sensors[j].SensorId)
	})

	unitDate.TimeSeries = data_models.TimeSeries{
		TimeDateStamp:  make([]time.Time, 0),
		TimeStamp:      make([]string, 0),
		TimeSeriesData: make([]data_models.TimeSeriesData, 0),
	}

	loc, _ := time.LoadLocation("Europe/London")

	for _, sensor := range unitDate.Sensors {
		if sensor.SensorId == "0" {
			for _, data := range sensor.TempData.Data {
				unitDate.TimeSeries.TimeDateStamp = append(unitDate.TimeSeries.TimeDateStamp, data.CreatedAt.Local())
				unitDate.TimeSeries.TimeStamp = append(unitDate.TimeSeries.TimeStamp, data.CreatedAt.In(loc).Format("15:00"))
			}
			break
		}
	}

	for index, sensor := range unitDate.Sensors {
		unitDate.TimeSeries.TimeSeriesData = append(unitDate.TimeSeries.TimeSeriesData, data_models.TimeSeriesData{
			SensorId:          sensor.SensorId,
			SensorName:        sensor.Name,
			SensorDescription: sensor.Description,
			TempReadingsInC:   make([]float32, 0),
		})

		for _, timeStamp := range unitDate.TimeSeries.TimeDateStamp {
			tempReading, err := l.dbClient.GetTempForSensor(ctx, id, sensor.SensorId, timeStamp)
			if err != nil {
				tempReading = -15
			}
			unitDate.TimeSeries.TimeSeriesData[index].TempReadingsInC = append(unitDate.TimeSeries.TimeSeriesData[index].TempReadingsInC, tempReading)
			unitDate.TimeSeries.TimeSeriesData[index].TempReadingsInPercentage = append(unitDate.TimeSeries.TimeSeriesData[index].TempReadingsInPercentage, float32(math.Max(float64(tempReading-5)/30.0, 0.0)))
		}
	}

	response, _ := json.Marshal(unitDate)
	res.StatusCode = http.StatusOK
	res.Body = string(response)
	return res, nil
}

func (l UnitLambdaHandler) insertUnitData(ctx context.Context, req events.APIGatewayProxyRequest, res handler.Response) (handler.Response, error) {

	var unit data_models.Unit
	if err := json.Unmarshal([]byte(req.Body), &unit); err != nil {
		log.Printf("error unmarshalling data - bad unit data %s", err.Error())
		lambdaResponse := LambdaResponse{
			Message: fmt.Sprintf("bad unit data"),
		}
		response, _ := json.Marshal(lambdaResponse)
		res.StatusCode = http.StatusBadRequest
		res.Body = string(response)
		return res, err
	}

	log.Printf("Calling InsertUnitData %+v", unit)

	if err := l.dbClient.InsertUnitData(ctx, &unit); err != nil {
		log.Printf("InsertUnitData error: %s", err.Error())
		lambdaResponse := LambdaResponse{
			Message: fmt.Sprintf("failed storing unit data"),
		}
		response, _ := json.Marshal(lambdaResponse)
		res.StatusCode = http.StatusInternalServerError
		res.Body = string(response)
		return res, err
	}
	res.StatusCode = http.StatusCreated
	return res, nil
}

func (l UnitLambdaHandler) updateUnit(ctx context.Context, req events.APIGatewayProxyRequest, res handler.Response) (handler.Response, error) {
	lambdaResponse := LambdaResponse{
		Message: fmt.Sprintf("Put unit"),
	}
	response, err := json.Marshal(lambdaResponse)

	res.StatusCode = http.StatusOK
	res.Body = string(response)
	return res, err
}
