package mysql

import (
	"api-gmr/store"
	"api-gmr/store/repository"
)

//MysqlDriver represent mysql driver
type MysqlDriver struct{}

//GetUserRepository returns mysql repository.User instance
func (d MysqlDriver) GetUserRepository() repository.User {
	return NewUserRepo()
}

func init() {
	store.Register("mysql", &MysqlDriver{})
}
