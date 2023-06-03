/*
* API Gateway
*/

resource "aws_api_gateway_rest_api" "homeApi" {
  name = local.name
}

resource "aws_api_gateway_resource" "homeApi" {
  rest_api_id = aws_api_gateway_rest_api.homeApi.id
  parent_id   = aws_api_gateway_rest_api.homeApi.root_resource_id
  path_part   = "hello"
}

resource "aws_api_gateway_method" "homeApi" {
  rest_api_id   = aws_api_gateway_rest_api.homeApi.id
  resource_id   = aws_api_gateway_resource.homeApi.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "homeApi" {
  rest_api_id = aws_api_gateway_rest_api.homeApi.id
  resource_id = aws_api_gateway_resource.homeApi.id
  http_method = aws_api_gateway_method.homeApi.http_method
  status_code = "200"
}

resource "aws_api_gateway_integration" "homeApi" {
  depends_on = [aws_api_gateway_method.homeApi, aws_api_gateway_method_response.homeApi]

  rest_api_id             = aws_api_gateway_rest_api.homeApi.id
  resource_id             = aws_api_gateway_method.homeApi.resource_id
  http_method             = aws_api_gateway_method.homeApi.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.func.invoke_arn
}

resource "aws_api_gateway_integration_response" "homeApi" {
  depends_on = [aws_api_gateway_integration.homeApi]

  rest_api_id = aws_api_gateway_rest_api.homeApi.id
  resource_id = aws_api_gateway_resource.homeApi.id
  http_method = aws_api_gateway_method.homeApi.http_method
  status_code = aws_api_gateway_method_response.homeApi.status_code

  response_templates = {
    "application/json" = ""
  }
}

module "cors" {
  source = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.homeApi.id
  api_resource_id = aws_api_gateway_resource.homeApi.id

  allow_headers = [
    "Authorization",
    "Content-Type",
    "X-Amz-Date",
    "X-Amz-Security-Token",
    "X-Api-Key",
    "X-Charge"
  ]
}

resource "aws_api_gateway_deployment" "homeApi" {
  depends_on = [aws_api_gateway_integration_response.homeApi]

  rest_api_id = aws_api_gateway_rest_api.homeApi.id
  description = "Deployed endpoint at ${timestamp()}"
}

resource "aws_api_gateway_stage" "homeApi" {
  stage_name    = local.environment
  rest_api_id   = aws_api_gateway_rest_api.homeApi.id
  deployment_id = aws_api_gateway_deployment.homeApi.id
}

resource "aws_lambda_permission" "homeApi" {
  statement_id  = "${local.name}-AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = local.name
  principal = "apigateway.amazonaws.com"
  source_arn = "arn:aws:execute-api:${local.region}:${local.account_id}:${aws_api_gateway_rest_api.homeApi.id}/*/${aws_api_gateway_method.homeApi.http_method}${aws_api_gateway_resource.homeApi.path}"
}
