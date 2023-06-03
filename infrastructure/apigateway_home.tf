#/*
#* API Gateway - home-api
#*/
#
#resource "aws_api_gateway_rest_api" "homeApi" {
#  name = local.homeApi
#}
#
#resource "aws_api_gateway_resource" "home" {
#  rest_api_id = aws_api_gateway_rest_api.homeApi.id
#  parent_id   = aws_api_gateway_rest_api.homeApi.root_resource_id
#  path_part   = "home"
#}
#
#resource "aws_api_gateway_method" "home" {
#  rest_api_id   = aws_api_gateway_rest_api.homeApi.id
#  resource_id   = aws_api_gateway_resource.home.id
#  http_method   = "GET"
#  authorization = "NONE"
#}
#
#resource "aws_api_gateway_method_response" "home" {
#  rest_api_id = aws_api_gateway_rest_api.homeApi.id
#  resource_id = aws_api_gateway_resource.home.id
#  http_method = aws_api_gateway_method.home.http_method
#  status_code = "200"
#}
#
#resource "aws_api_gateway_integration" "home" {
#  depends_on = [aws_api_gateway_method.home, aws_api_gateway_method_response.home]
#
#  rest_api_id             = aws_api_gateway_rest_api.homeApi.id
#  resource_id             = aws_api_gateway_method.home.resource_id
#  http_method             = aws_api_gateway_method.home.http_method
#  integration_http_method = "POST"
#  type                    = "AWS_PROXY"
#  uri                     = aws_lambda_function.func.invoke_arn
#}
#
#resource "aws_api_gateway_integration_response" "home" {
#  depends_on = [aws_api_gateway_integration.home]
#
#  rest_api_id = aws_api_gateway_rest_api.homeApi.id
#  resource_id = aws_api_gateway_resource.home.id
#  http_method = aws_api_gateway_method.home.http_method
#  status_code = aws_api_gateway_method_response.home.status_code
#
#  response_templates = {
#    "application/json" = ""
#  }
#}
#
#module "homeCors" {
#  source = "squidfunk/api-gateway-enable-cors/aws"
#  version = "0.3.3"
#
#  api_id          = aws_api_gateway_rest_api.homeApi.id
#  api_resource_id = aws_api_gateway_resource.home.id
#
#  allow_headers = [
#    "Authorization",
#    "Content-Type",
#    "X-Amz-Date",
#    "X-Amz-Security-Token",
#    "X-Api-Key",
#    "X-Charge"
#  ]
#}
#
#resource "aws_api_gateway_deployment" "homeApi" {
#  depends_on = [aws_api_gateway_integration_response.home]
#
#  rest_api_id = aws_api_gateway_rest_api.homeApi.id
#  description = "Deployed endpoint at ${timestamp()}"
#}
#
#resource "aws_api_gateway_stage" "homeApi" {
#  stage_name    = local.environment
#  rest_api_id   = aws_api_gateway_rest_api.homeApi.id
#  deployment_id = aws_api_gateway_deployment.homeApi.id
#}
#
#resource "aws_lambda_permission" "homeApi" {
#  statement_id  = "${local.home}-AllowExecutionFromAPIGateway"
#  action = "lambda:InvokeFunction"
#  function_name = local.name
#  principal = "apigateway.amazonaws.com"
#  source_arn = "arn:aws:execute-api:${local.region}:${local.account_id}:${aws_api_gateway_rest_api.homeApi.id}/*/${aws_api_gateway_method.home.http_method}${aws_api_gateway_resource.home.path}"
#}
