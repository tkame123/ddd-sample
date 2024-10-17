package repository

import (
	"context"
)

type ProcessedMessage interface {
	ProcessedMessageExists(ctx context.Context, messageID string) (bool, error)
	ProcessedMessageSave(ctx context.Context, messageID string) error
}
