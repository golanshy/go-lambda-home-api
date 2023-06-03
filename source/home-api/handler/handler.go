package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

// Handler - interface
type Handler interface {
	HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error)
}
