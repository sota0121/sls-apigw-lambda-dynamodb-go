package main

import (
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

func GetUserHandler(request GetUserRequest) (GetUserResponse, error) {
	return GetUserResponse{
		UserId:      request.UserId,
		UserName:    "test_user",
		CompanyName: "test_company",
		Authority:   "admin",
		GoogleId:    "test_google_id",
	}, nil
}

func main() {
	lambda.Start(GetUserHandler)
}
