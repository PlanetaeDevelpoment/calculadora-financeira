package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/EdmilsonRodrigues/calculadora-financeira/calculadora-recisao/calculadora"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}



func handler(ctx context.Context, event calculadorarecisao.RequisiçãoRecisão) (events.APIGatewayProxyResponse, error) {
	log.Println("Received event:", event)
	result := calculadorarecisao.CalcularRecisão(event)
    jsonResult, err := json.Marshal(result)
    if err != nil {
		log.Println("Error marshalling JSON:", err)
        return events.APIGatewayProxyResponse{
            Body:       "Error marshalling JSON",
            StatusCode: 500,
        }, nil
    }
    return events.APIGatewayProxyResponse{
        Body:       string(jsonResult),
        StatusCode: 200,
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
    }, nil
}
