terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.1.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5.1"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.3.0"
    }
  }

  required_version = "~> 1.0"
}


provider "aws" {
  region = "eu-west-2"
  profile = "applylogic"
}

data "aws_caller_identity" "current" {}

locals {
  account_id          = data.aws_caller_identity.current.account_id
  environment         = "dev"
  lambda_handler      = "home-api"
  name                = "go-lambda-home-api"
  hello_api           = "hello-api"
  home_api            = "home-api"
  hello_message       = "Hello World"
  region              = "eu-west-2"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../bin/home-api"
  output_path = "bin/home-api.zip"
}

