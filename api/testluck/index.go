package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
	"github.com/lambda-direct/gocast-wager/dbclient"
	"github.com/lambda-direct/gocast-wager/env"

	"github.com/lambda-direct/gocast-wager/flipcoinclient"
)

type RequestBody struct {
	Username string `json:"username" required:"true"`
	Amount   uint16 `json:"amount" required:"true"`
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var s env.Spec
	if err := envconfig.Process("", &s); err != nil {
		return nil, err
	}

	var reqBody RequestBody
	if err := json.Unmarshal([]byte(req.Body), &reqBody); err != nil {
		return nil, fmt.Errorf("unable to decode request body: %w", err)
	}

	fcClient, err := flipcoinclient.New(os.Getenv("API_HOST"))
	if err != nil {
		return nil, err
	}

	dbClient, err := dbclient.New(s.DB)
	if err != nil {
		return nil, err
	}

	defer dbClient.Close()

	luck, err := fcClient.TestLuck()
	if err != nil {
		return nil, err
	}

	flipResult := &dbclient.FlipResult{
		Username: reqBody.Username,
		Amount:   reqBody.Amount,
		Result:   luck,
	}

	err = dbClient.SaveResult(flipResult)
	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(flipResult)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(res),
	}, nil
}

func main() {
	lambda.Start(handler)
}
