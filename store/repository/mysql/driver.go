package mysql

import (
	"api-gmr/store"
	"api-gmr/store/repository"
)

type Driver struct{}

func (d Driver) GetUserRepository() repository.User {
	return NewUserRepo()
}

func init() {
	store.Register("mysql", &Driver{})
}
