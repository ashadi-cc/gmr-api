package repository

type UserModel interface {
	GetUserID() int
	GetUsername() string
	GetEmail() string
	GetPasswordHash() string
	GetUserGroup() string
	GetBlok() string
	getName() string
}

type User interface {
	FindByUsername(username string) (UserModel, error)
	UpdateEmailandPassword(user UserModel) error
}
