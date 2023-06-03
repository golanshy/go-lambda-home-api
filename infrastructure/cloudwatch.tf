/*
* Cloudwatch
*/

// Log group

resource "aws_cloudwatch_log_group" "main_api_gw" {
  name = "/aws/api-gw/${aws_apigatewayv2_api.home_api.name}"

  retention_in_days = 14
}