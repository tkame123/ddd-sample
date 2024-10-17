package repository

import (
	"context"
)

type ProcessedMessage interface {
	ProcessedMessageSave(ctx context.Context, messageID string) error
	ProcessedMessageDelete(ctx context.Context, messageID string) error
}
