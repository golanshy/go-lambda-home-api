terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.46.0"
    }
  }
}

provider "aws" {
  region = "eu-west-2"
}

data "aws_caller_identity" "current" {}

locals {
  account_id     = data.aws_caller_identity.current.account_id
  environment    = "dev"
  lambda_handler = "home-api"
  name           = "go-lambda-home-api"
  random_name    = "Hello World"
  region         = "eu-west-2"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../bin/home-api"
  output_path = "bin/home-api.zip"
}
