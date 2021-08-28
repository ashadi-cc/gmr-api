package provider

import "api-gmr/store/repository"

//Driver base Driver methods interface
type Driver interface {
	//GetUserRepository returns a new repository.User instance
	GetUserRepository() repository.User

	//GetBillingRepository returns a new repository.User instance
	GetBillingRepository() repository.Billing
}
