package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/processedmessage"
)

func (r *repo) ProcessedMessageExists(ctx context.Context, messageID string) (bool, error) {
	return r.db.ProcessedMessage.Query().
		Where(processedmessage.MessageID(messageID)).
		Exist(ctx)
}

func (r *repo) ProcessedMessageSave(ctx context.Context, messageID string) error {
	_, err := r.db.ProcessedMessage.Create().
		SetMessageID(messageID).
		Save(ctx)
	return err
}
