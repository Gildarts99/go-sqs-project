package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMClient struct {
	client *ssm.Client
}

func NewSSMClient() *SSMClient {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := ssm.NewFromConfig(cfg)
	return &SSMClient{
		client: client,
	}
}

func (s *SSMClient) GetParameter(path string, secureString bool) (string, error) {
	parameter, err := s.client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           &path,
		WithDecryption: &secureString,
	})

	if err != nil {
		return "", err
	}

	return *parameter.Parameter.Value, nil
}
