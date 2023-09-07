package main

import (
	"context"

	"github.com/Asad2730/IAM_Example/services"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	client := iam.NewFromConfig(cfg)

	services.CreateUser(client)
}
