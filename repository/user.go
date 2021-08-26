package repository

import "context"

type UserModel interface {
	GetUserID() int
	GetUsername() string
	GetEmail() string
	GetPasswordHash() string
	GetUserGroup() string
	GetBlok() string
	GetName() string
}

type User interface {
	FindByUsername(ctx context.Context, username string) (UserModel, error)
	FindByUserID(ctx context.Context, userID int) (UserModel, error)
	UpdateEmailandPassword(ctx context.Context, user UserModel) error
}
