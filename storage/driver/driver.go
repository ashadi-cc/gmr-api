package driver

import "api-gmr/storage/repo"

//Driver base file driver methods
type Driver interface {
	repo.FileRepo
	//GetName returns driver name
	GetName() string
}
