package idempotency

import "errors"

type NotFoundError struct {
	id string
}

func (e *NotFoundError) Error() string {
	return "dynamodb: " + e.id + " not found"
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}
