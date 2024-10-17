package database

import (
	"context"
	"log"
)

func (r repo) ProcessedMessageExists(ctx context.Context, messageID string) (bool, error) {
	//TODO implement me
	log.Println("implement me: ProcessedMessageExists")
	return false, nil
}

func (r repo) ProcessedMessageSave(ctx context.Context, messageID string) error {
	//TODO implement me
	log.Println("implement me: ProcessedMessageSave")
	return nil
}
