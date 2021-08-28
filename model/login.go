package model

//UserLogin to hold user login input values
type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//Validate validate UserLogin input values
func (user UserLogin) Validate() error {
	return validate.Struct(user)
}
