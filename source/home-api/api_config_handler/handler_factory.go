package api_config_handler

type ConfigLambdaHandler struct {
}

// Create -
func Create() *ConfigLambdaHandler {
	return NewLambdaHandler()
}

// NewLambdaHandler -
func NewLambdaHandler() *ConfigLambdaHandler {
	return &ConfigLambdaHandler{}
}
