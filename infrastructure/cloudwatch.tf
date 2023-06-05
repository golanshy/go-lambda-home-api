/*
* Cloudwatch
*/

// Log group

resource "aws_cloudwatch_log_group" "main_api_gw" {
#  name = "/aws/lambda/${aws_apigatewayv2_api.home_api.name}"
  name = "/aws/lambda/go-lambda-api"
  retention_in_days = 14
}
