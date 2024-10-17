package database

import (
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/repository"
)

type repo struct {
}

func NewRepository() repository.Repository {
	return &repo{}
}
