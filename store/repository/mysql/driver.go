package mysql

import (
	"api-gmr/store"
	"api-gmr/store/repository"
)

type MysqlDriver struct{}

func (d MysqlDriver) GetUserRepository() repository.User {
	return NewUserRepo()
}

func init() {
	store.Register("mysql", &MysqlDriver{})
}
