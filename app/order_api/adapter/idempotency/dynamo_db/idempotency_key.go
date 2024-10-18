package dynamo_db

import "time"

const TableName = "idempotency_key"
const TtlDuration = 60 * time.Minute

// MEMO: TTLに対しての自動削除はDynamoDB側で対応可能な模様
// https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/TTL.html
type IdempotencyKey struct {
	ID   IdempotencyKeyID `dynamodbav:"id"`   // PK
	Type TypeName         `dynamodbav:"type"` // SK
	TTL  int64            `dynamodbav:"ttl"`
	Data string           `dynamodbav:"data"`
}

type IdempotencyKeyID = string
type TypeName = string

const (
	TypeNameRequest  = "Request"
	TypeNameResponse = "Response"
	TypeNameStatus   = "Status"
)
