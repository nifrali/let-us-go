package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/nifrali/let-us-go/sqltest"
)

func main() {
	lambda.Start(connHandler)
}

func connHandler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	version, err := sqltest.TestDbConnection(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("DATABASE"),
		os.Getenv("INSTANCE"),
		os.Getenv("PORT"),
	)

	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       version,
	}, nil

}
