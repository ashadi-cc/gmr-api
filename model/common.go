package model

import "github.com/go-playground/validator/v10"

type CommonMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

func (c CommonMessage) WithError(err error) CommonMessage {
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs.Translate(trans) {
				c.Errors = append(c.Errors, e)
			}
		}
	}
	return c
}
