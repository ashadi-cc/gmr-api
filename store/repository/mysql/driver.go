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

//GetBillingRepository returns mysql repository.Billing instance
func (d MysqlDriver) GetBillingRepository() repository.Billing {
	return NewBillingRepo()
}

func init() {
	store.Register("mysql", &MysqlDriver{})
}
