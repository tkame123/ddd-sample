package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/processedmessage"
)

func (r *repo) ProcessedMessageDelete(ctx context.Context, messageID string) error {
	_, err := r.db.ProcessedMessage.Delete().
		Where(processedmessage.MessageID(messageID)).
		Exec(ctx)
	return err
}

func (r *repo) ProcessedMessageSave(ctx context.Context, messageID string) error {
	_, err := r.db.ProcessedMessage.Create().
		SetMessageID(messageID).
		Save(ctx)
	return err
}
