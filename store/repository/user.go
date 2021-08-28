package repository

import "context"

//UserModel base user methods collections
type UserModel interface {
	//GetUserID returns user id
	GetUserID() int
	//GetUsername returns user name
	GetUsername() string
	//GetEmail returns user email
	GetEmail() string
	//GetPasswordHash returns user hash password
	GetPasswordHash() string
	//GetGroup returns user group
	GetGroup() string
	//GetBlok returns user blok
	GetBlok() string
	//GetName returns user name
	GetName() string
}

//User represents User repository methods collection
type User interface {
	//FindByUsername returns UserModel by given username
	FindByUsername(ctx context.Context, username string) (UserModel, error)
	//FindByUserID returns UserModel by given userid
	FindByUserID(ctx context.Context, userID int) (UserModel, error)
	//UpdateEmailandPassword update user email and password by given user payload
	UpdateEmailandPassword(ctx context.Context, user UserModel) error
}
