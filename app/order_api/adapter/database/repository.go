package database

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type repo struct {
	db *ent.Client
}

func NewRepository(db *ent.Client) repository.Repository {
	return &repo{db: db}
}

func (r *repo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
