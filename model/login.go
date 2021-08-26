package model

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (user UserLogin) Validate() error {
	return validate.Struct(user)
}
