package api_unit_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/go-lambda-home-api/data_models"
	"github.com/golanshy/go-lambda-home-api/handler"
	"log"
	"net/http"
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

	//unitDate := &data_models.Unit{
	//	API:         data.API,
	//	Id:          id,
	//	Name:        data.Name,
	//	Description: data.Description,
	//	HomeId:      data.HomeId,
	//	Sensors:     data.Sensors,
	//	CreatedAt:   data.CreatedAt,
	//	UpdatedAt:   data.UpdatedAt,
	//}

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
