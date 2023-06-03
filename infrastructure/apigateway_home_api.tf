/*
* API Gateway
*/

resource "aws_api_gateway_resource" "homeAPi" {
  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  parent_id   = aws_api_gateway_rest_api.helloApi.root_resource_id
  path_part   = "home"
}

resource "aws_api_gateway_method" "homeAPi" {
  rest_api_id   = aws_api_gateway_rest_api.helloApi.id
  resource_id   = aws_api_gateway_resource.homeAPi.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "homeAPi" {
  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  resource_id = aws_api_gateway_resource.homeAPi.id
  http_method = aws_api_gateway_method.homeAPi.http_method
  status_code = "200"
}

resource "aws_api_gateway_integration" "homeAPi" {
  depends_on = [aws_api_gateway_method.homeAPi, aws_api_gateway_method_response.homeAPi]

  rest_api_id             = aws_api_gateway_rest_api.helloApi.id
  resource_id             = aws_api_gateway_method.homeAPi.resource_id
  http_method             = aws_api_gateway_method.homeAPi.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.func.invoke_arn
}

resource "aws_api_gateway_integration_response" "homeAPi" {
  depends_on = [aws_api_gateway_integration.homeAPi]

  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  resource_id = aws_api_gateway_resource.homeAPi.id
  http_method = aws_api_gateway_method.homeAPi.http_method
  status_code = aws_api_gateway_method_response.homeAPi.status_code

  response_templates = {
    "application/json" = ""
  }
}

module "cors-home-api" {
  source = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.helloApi.id
  api_resource_id = aws_api_gateway_resource.homeAPi.id

  allow_headers = [
    "Authorization",
    "Content-Type",
    "X-Amz-Date",
    "X-Amz-Security-Token",
    "X-Api-Key",
    "X-Charge"
  ]
}

resource "aws_api_gateway_deployment" "homeAPi" {
  depends_on = [aws_api_gateway_integration_response.homeAPi]

  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  description = "Deployed endpoint at ${timestamp()}"
}

resource "aws_api_gateway_stage" "homeAPi" {
  stage_name    = "${local.environment}-${local.home_api}"
  rest_api_id   = aws_api_gateway_rest_api.helloApi.id
  deployment_id = aws_api_gateway_deployment.homeAPi.id
}

resource "aws_lambda_permission" "homeAPi" {
  statement_id  = "${local.name}-${local.home_api}-AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = local.name
  principal = "apigateway.amazonaws.com"
  source_arn = "arn:aws:execute-api:${local.region}:${local.account_id}:${aws_api_gateway_rest_api.helloApi.id}/*/${aws_api_gateway_method.homeAPi.http_method}${aws_api_gateway_resource.homeAPi.path}"
}
