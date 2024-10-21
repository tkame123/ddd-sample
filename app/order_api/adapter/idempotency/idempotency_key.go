package idempotency

import "connectrpc.com/connect"

type IdempotencyKey struct {
	ID       IdempotencyKeyID
	Status   Status
	Request  connect.AnyRequest
	Response connect.AnyResponse
}

type IdempotencyKeyID = string
type Status = string

const (
	IdempotencyKeyStatusProcessing Status = "Processing"
	IdempotencyKeyStatusSuccess    Status = "Success"
)
