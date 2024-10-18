package Idempotency_key

type IdempotencyKey struct {
	ID       IdempotencyKeyID
	Status   Status
	Request  any
	Response any
}

type IdempotencyKeyID = string
type Status = string

const (
	IdempotencyKeyStatusProcessing Status = "Processing"
	IdempotencyKeyStatusSuccess    Status = "Success"
)
