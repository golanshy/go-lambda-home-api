/*
* API Gateway
*/

resource "aws_apigatewayv2_integration" "lambda_home_api" {
  api_id = aws_apigatewayv2_api.home_api.id

  integration_uri    = aws_lambda_function.func.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "get_hello" {
  api_id = aws_apigatewayv2_api.home_api.id

  route_key = "GET /hello"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_home_api.id}"
}

resource "aws_apigatewayv2_route" "get_home" {
  api_id = aws_apigatewayv2_api.home_api.id

  route_key = "GET /home"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_home_api.id}"
}

resource "aws_apigatewayv2_route" "post_hello" {
  api_id = aws_apigatewayv2_api.home_api.id

  route_key = "POST /home"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_home_api.id}"
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.func.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_api.execution_arn}/*/*"
}

