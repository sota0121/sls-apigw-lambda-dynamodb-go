package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type GetUserRequest struct {
	UserId string `json:"user_id"`
}

type GetUserResponse struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	CompanyName string `json:"company_name"`
	Authority   string `json:"authority"`
	GoogleId    string `json:"google_id"`
}

func GetUserHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryStringParameters := req.QueryStringParameters
	fmt.Println(queryStringParameters)
	pathParameters := req.PathParameters
	fmt.Println(pathParameters)
	headers := req.Headers
	fmt.Println(headers)
	body := req.Body
	fmt.Println(body)

	userId := pathParameters["id"]
	fmt.Println("path parameter: id = " + userId)

	respBody := GetUserResponse{
		UserId:      "test_user_id",
		UserName:    "test_user",
		CompanyName: "test_company",
		Authority:   "admin",
		GoogleId:    "test_google_id",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%v", respBody),
	}, nil
}

func main() {
	lambda.Start(GetUserHandler)
}
