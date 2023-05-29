/*
* Cloudwatch
*/

// Log group
resource "aws_cloudwatch_log_group" "log" {
  name              = "/aws/lambda/${local.name}"
  retention_in_days = 7
}

resource "aws_cloudwatch_log_group" "homeApiLog" {
  name              = "/aws/lambda/${local.homeApi}"
  retention_in_days = 7
}
