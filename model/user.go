package model

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Group    string `json:"-"`
	Username string `json:"username"`
	Blok     string `json:"blok"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (user User) GetUserID() int {
	return user.Id
}

func (user User) GetEmail() string {
	return user.Email
}

func (user User) GetGroup() string {
	return user.Group
}

func (user User) GetPasswordHash() string {
	return user.Password
}

func (user User) GetBlok() string {
	return user.Blok
}

func (user User) GetUsername() string {
	return user.Username
}

func (user User) GetName() string {
	return user.Name
}

func (user *User) SetUserId(id int) {
	user.Id = id
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user *User) SetGroup(group string) {
	user.Group = group
}

type UserInput struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
}

func (user UserInput) Validate() error {
	return validate.Struct(user)
}
