package idempotency

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/idempotency/dynamo_db"
	"time"
)

type Repository struct {
	dynamoClient *dynamodb.Client
}

func NewRepository(dynamoClient *dynamodb.Client) *Repository {
	return &Repository{dynamoClient: dynamoClient}
}

func (r *Repository) IdempotencyKeyFindByID(ctx context.Context, id string) (*IdempotencyKey, error) {
	response, err := r.dynamoClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(dynamo_db.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
			"type": &types.AttributeValueMemberS{
				Value: dynamo_db.TypeNameStatus,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(response.Item) == 0 {
		return nil, &NotFoundError{id: id}
	}

	var res dynamo_db.IdempotencyKey
	err = attributevalue.UnmarshalMap(response.Item, &res)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response. Here's why: %w\n", err)
	}

	return ToModel(&res)
}

func (r *Repository) IdempotencyKeySave(ctx context.Context, key IdempotencyKey) error {
	var writeReqs []types.WriteRequest

	rawItems := ToDynamoDB(key)
	for _, v := range rawItems {
		item, err := attributevalue.MarshalMap(v)
		if err != nil {
			return err
		}
		writeReqs = append(
			writeReqs,
			types.WriteRequest{PutRequest: &types.PutRequest{Item: item}},
		)
	}
	_, err := r.dynamoClient.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{dynamo_db.TableName: writeReqs}})
	if err != nil {
		return err
	}

	return nil
}

func ToDynamoDB(m IdempotencyKey) []*dynamo_db.IdempotencyKey {
	ttl := newTTL().Unix()
	data := make([]*dynamo_db.IdempotencyKey, 0, 3)

	data = append(data, &dynamo_db.IdempotencyKey{
		ID:   m.ID,
		Type: dynamo_db.TypeNameRequest,
		TTL:  ttl,
		Data: fmt.Sprintf("%s", m.Request.Any()),
	})
	data = append(data, &dynamo_db.IdempotencyKey{
		ID:   m.ID,
		Type: dynamo_db.TypeNameStatus,
		TTL:  ttl,
		Data: m.Status,
	})

	if m.Response != nil {
		data = append(data, &dynamo_db.IdempotencyKey{
			ID:   m.ID,
			Type: dynamo_db.TypeNameResponse,
			TTL:  ttl,
			Data: fmt.Sprintf("%s", m.Response.Any()),
		})
	}

	return data
}

func ToModel(r *dynamo_db.IdempotencyKey) (*IdempotencyKey, error) {
	// MEMO: 一旦Statusだけ返している
	if r.Type != dynamo_db.TypeNameStatus {
		return nil, errors.New("Type is not Status")
	}
	return &IdempotencyKey{
		ID:     r.ID,
		Status: r.Data,
	}, nil
}

func newTTL() time.Time {
	return time.Now().Add(dynamo_db.TtlDuration)
}
