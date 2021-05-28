package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(2)
	randBool := randInt != 0

	var res string
	if randBool {
		res = "Heads"
	} else {
		res = "Tails"
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("\"%s\"", res),
	}, nil
}

func main() {
	lambda.Start(handler)
}
