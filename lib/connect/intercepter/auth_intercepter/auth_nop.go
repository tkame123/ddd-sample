package auth_intercepter

import (
	"context"
	"github.com/tkame123/ddd-sample/lib/auth"
	"github.com/tkame123/ddd-sample/lib/metadata"
	"log"
)

type authStrategyNop struct {
}

func NewAuthStrategyNop() auth.Strategy {
	return &authStrategyNop{}
}

func (a *authStrategyNop) Authenticate(ctx context.Context) (*metadata.UserInfo, error) {
	log.Println("nop Authn")

	return &metadata.UserInfo{
		ID: "develop user",
	}, nil
}

func (a *authStrategyNop) Authorize(ctx context.Context) error {
	log.Println("nop Authz")

	return nil
}
