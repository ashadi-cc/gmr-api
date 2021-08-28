package model

//User to hold User data
type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Group    string `json:"-"`
	Username string `json:"username"`
	Blok     string `json:"blok"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

//GetUserID returns User.id
func (user User) GetUserID() int {
	return user.Id
}

//GetEmail returns User.Email
func (user User) GetEmail() string {
	return user.Email
}

//GetGroup returns User.Group
func (user User) GetGroup() string {
	return user.Group
}

//GetPasswordHash returns user.Password
func (user User) GetPasswordHash() string {
	return user.Password
}

//GetBlok returns user.Blok
func (user User) GetBlok() string {
	return user.Blok
}

//GetUsername returns User.Username
func (user User) GetUsername() string {
	return user.Username
}

//GetName returns User.Name
func (user User) GetName() string {
	return user.Name
}

//SetUserId set User.Id by given id
func (user *User) SetUserId(id int) {
	user.Id = id
}

//SetEmail set User.Email gy given email
func (user *User) SetEmail(email string) {
	user.Email = email
}

//SetGroup set User.Group by given group
func (user *User) SetGroup(group string) {
	user.Group = group
}

//UserInput to hold user data
type UserInput struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
}

//Validate validate UserInput values
func (user UserInput) Validate() error {
	return validate.Struct(user)
}
