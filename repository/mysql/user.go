package mysql

import "api-gmr/repository"

type userRepo struct{}

func NewUserRepo() repository.User {
	return &userRepo{}
}

func (repo userRepo) FindByUsername(username string) (repository.UserModel, error) {
	return nil, nil
}

func (repo userRepo) UpdateEmailandPassword(user repository.UserModel) error {
	return nil
}
