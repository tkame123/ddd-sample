package domain_event

import "context"

type DomainEventHandle[T any] interface {
	Handler(ctx context.Context, event T) error
}
