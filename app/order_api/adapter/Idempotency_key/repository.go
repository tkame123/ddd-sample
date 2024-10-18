package Idempotency_key

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Repository struct {
	dynamoClient *dynamodb.Client
}

func NewRepository(dynamoClient *dynamodb.Client) *Repository {
	return &Repository{dynamoClient: dynamoClient}
}

func (r *Repository) IdempotencyKeyFindByID(ctx context.Context, id string) (*IdempotencyKey, error) {
	// TODO: implement
	panic("implement me")
	return nil, nil
}

func (r *Repository) IdempotencyKeySave(ctx context.Context, key IdempotencyKey) error {

	panic("implement me")
}
