package metadata

import (
	"context"
)

type UserID = string
type Permission = string

type UserInfo struct {
	ID           UserID
	AccessPolicy *AccessPolicy
}

type AccessPolicy struct {
	Permissions []Permission
}

type userInfoKey struct{}

func WithUserInfo(ctx context.Context, user *UserInfo) context.Context {
	return context.WithValue(ctx, userInfoKey{}, user)
}

func GetUserInfo(ctx context.Context) (*UserInfo, bool) {
	user, ok := ctx.Value(userInfoKey{}).(*UserInfo)
	return user, ok
}
