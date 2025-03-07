#!/bin/bash

if ! GOOS=linux GOARCH=amd64 go build -o bootstrap main.go; then
  echo "Build failed"
  exit 1
fi

if ! zip deployment.zip bootstrap; then
  echo "Zip failed"
  exit 1
fi

if ! aws lambda update-function-code --function-name calculadora-recisao --zip-file fileb://deployment.zip; then
  echo "Update failed"
  exit 1
fi
