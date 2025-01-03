/*
* Cloudwatch
*/

// Log group

resource "aws_cloudwatch_log_group" "main_api_gw" {
  name = "/aws/lambda/lambda-home-api"
  retention_in_days = 1
}
