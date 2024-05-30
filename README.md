# go-lambda-terraform-setup

Setup a Go Lambda function with Terraform.

## DSTFix - Day light saving time fix
Update the request_handler.go file DSTFix value:
DSTFix = 0 - no fix for DST
DSTFix = 1 - Add daylight Saving time 1 hour

`configResponse := ConfigResponse{
    EnableLocalLogs: false,
    DelayInSeconds:  1 * 60 * 60, // 1 hour
    DSTFix:          1,           // Daylight Saving time = 1 hour
}`


## Setup

`chmod +x ./deploy.sh`

`chmod +x ./destroy.sh`

## Usage

- To deploy run `sh deploy.sh`

- To destroy run `sh destroy.sh`
