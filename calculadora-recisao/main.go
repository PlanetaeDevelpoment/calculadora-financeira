package main

import (
	"encoding/json"
	"fmt"

	"github.com/EdmilsonRodrigues/calculadora-financeira/calculadora-recisao/calculadora"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requisicao calculadorarecisao.RequisiçãoRecisão
	err := json.Unmarshal([]byte(request.Body), &requisicao)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid JSON payload",
			StatusCode: 400,
		}, nil
	}

	result := calculadorarecisao.CalcularRecisão(requisicao)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", result),
		StatusCode: 200,
	}, nil
}
