/*
* Lambda
*/

// Function
resource "aws_lambda_function" "func" {
  filename         = data.archive_file.lambda_zip.output_path
  function_name    = local.name
  role             = aws_iam_role.lambda.arn
  handler          = local.lambda_handler
  source_code_hash = filebase64sha256(data.archive_file.lambda_zip.output_path)
  runtime          = "go1.x"
  memory_size      = 1024
  timeout          = 30

  environment {
    variables = {
      HELLO_MESSAGE = local.hello_message
      SECRETS       = data.aws_secretsmanager_secret_version.secret_credentials.secret_string
    }
  }
}
