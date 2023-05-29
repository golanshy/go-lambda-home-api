/*
* API Gateway
*/

resource "aws_api_gateway_rest_api" "helloApi" {
  name = local.name
}

resource "aws_api_gateway_resource" "hello" {
  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  parent_id   = aws_api_gateway_rest_api.helloApi.root_resource_id
  path_part   = "hello"
}

resource "aws_api_gateway_method" "hello" {
  rest_api_id   = aws_api_gateway_rest_api.helloApi.id
  resource_id   = aws_api_gateway_resource.hello.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "hello" {
  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  resource_id = aws_api_gateway_resource.hello.id
  http_method = aws_api_gateway_method.hello.http_method
  status_code = "200"
}

resource "aws_api_gateway_integration" "hello" {
  depends_on = [aws_api_gateway_method.hello, aws_api_gateway_method_response.hello]

  rest_api_id             = aws_api_gateway_rest_api.helloApi.id
  resource_id             = aws_api_gateway_method.hello.resource_id
  http_method             = aws_api_gateway_method.hello.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.func.invoke_arn
}

resource "aws_api_gateway_integration_response" "hello" {
  depends_on = [aws_api_gateway_integration.hello]

  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  resource_id = aws_api_gateway_resource.hello.id
  http_method = aws_api_gateway_method.hello.http_method
  status_code = aws_api_gateway_method_response.hello.status_code

  response_templates = {
    "application/json" = ""
  }
}

module "cors" {
  source = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.helloApi.id
  api_resource_id = aws_api_gateway_resource.hello.id

  allow_headers = [
    "Authorization",
    "Content-Type",
    "X-Amz-Date",
    "X-Amz-Security-Token",
    "X-Api-Key",
    "X-Charge"
  ]
}

resource "aws_api_gateway_deployment" "helloApi" {
  depends_on = [aws_api_gateway_integration_response.hello]

  rest_api_id = aws_api_gateway_rest_api.helloApi.id
  description = "Deployed endpoint at ${timestamp()}"
}

resource "aws_api_gateway_stage" "helloApi" {
  stage_name    = local.environment
  rest_api_id   = aws_api_gateway_rest_api.helloApi.id
  deployment_id = aws_api_gateway_deployment.helloApi.id
}

resource "aws_lambda_permission" "helloApi" {
  statement_id  = "${local.name}-AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = local.name
  principal = "apigateway.amazonaws.com"
  source_arn = "arn:aws:execute-api:${local.region}:${local.account_id}:${aws_api_gateway_rest_api.helloApi.id}/*/${aws_api_gateway_method.hello.http_method}${aws_api_gateway_resource.hello.path}"
}
