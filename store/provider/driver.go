package provider

import "api-gmr/repository"

type Driver interface {
	GetUserRepository() repository.User
}
