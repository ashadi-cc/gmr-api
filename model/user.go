package model

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Group    string `json:"-"`
	Username string `json:"username"`
	Blok     string `json:"blok"`
	Name     string `json:"name"`
}

func (user User) GetUserId() int {
	return user.Id
}

func (user User) GetEmail() string {
	return user.Email
}

func (user User) GetGroup() string {
	return user.Group
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
