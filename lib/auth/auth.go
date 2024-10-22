package auth

import (
	"context"
)

type Strategy interface {
	Authenticate(ctx context.Context) error
	Authorize(ctx context.Context) error
}

type Context struct {
	strategy Strategy
}

func (a *Context) SetAuthStrategy(strategy Strategy) {
	a.strategy = strategy
}

func (a *Context) Authenticate(ctx context.Context) error {
	return a.strategy.Authenticate(ctx)
}

func (a *Context) Authorize(ctx context.Context) error {
	return a.strategy.Authorize(ctx)
}
