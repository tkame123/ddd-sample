package database

import (
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type repo struct {
	db *ent.Client
}

func NewRepository(db *ent.Client) repository.Repository {
	return &repo{db: db}
}
