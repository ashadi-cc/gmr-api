package mysql

import (
	"api-gmr/repository"
	"api-gmr/store"
)

type Driver struct{}

func (d Driver) GetUserRepository() repository.User {
	return NewUserRepo()
}

func init() {
	store.Register("mysql", &Driver{})
}
