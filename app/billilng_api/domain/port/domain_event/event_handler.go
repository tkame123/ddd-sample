package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/proto/message"
)

type EventHandler interface {
	Handler(ctx context.Context, mes *message.Message) error
}