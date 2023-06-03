/*
* Cloudwatch
*/

// Log group
resource "aws_cloudwatch_log_group" "log" {
  name = "/aws/api-gw/${aws_apigatewayv2_api.home_api.name}"
  retention_in_days = 14
}

