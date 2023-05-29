#!/bin/bash

echo -e "\n+++++ Starting deployment +++++\n"

tfswitch --show-latest
#tfswitch 1.0.0

rm -rf ./bin

echo "+++++ build go packages +++++"

cd source/home-api
go test ./...
env GOOS=linux GOARCH=amd64 go build -o ../../bin/home-api
cd ../..

echo "+++++ hello module +++++"
cd infrastructure
terraform init -input=false
terraform apply -input=false -auto-approve
cd ../

echo -e "\n+++++ Deployment done +++++\n"