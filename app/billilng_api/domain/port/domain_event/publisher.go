package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/proto/message"
)

type Publisher interface {
	PublishMessages(ctx context.Context, events []*message.Message)
}
