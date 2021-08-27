package provider

import "api-gmr/store/repository"

type Driver interface {
	GetUserRepository() repository.User
}
