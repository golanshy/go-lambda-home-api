package api_hello_handler

type HelloLambdaHandler struct {
}

// Create -
func Create() *HelloLambdaHandler {
	return NewLambdaHandler()
}

// NewLambdaHandler -
func NewLambdaHandler() *HelloLambdaHandler {
	return &HelloLambdaHandler{}
}
