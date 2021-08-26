package mysql

type user struct {
	id           int
	username     string
	email        string
	passwordHash string
	userGroup    string
	blok         string
	name         string
}

func (u user) GetUserID() int {
	return u.id
}

func (u user) GetUsername() string {
	return u.username
}

func (u user) GetEmail() string {
	return u.email
}

func (u user) GetPasswordHash() string {
	return u.passwordHash
}

func (u user) GetUserGroup() string {
	return u.userGroup
}

func (u user) GetBlok() string {
	return u.blok
}

func (u user) GetName() string {
	return u.name
}
