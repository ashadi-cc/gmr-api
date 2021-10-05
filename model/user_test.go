package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInput_Validate(t *testing.T) {
	type fields struct {
		Email    string
		Password string
	}
	tests := []struct {
		name      string
		fields    fields
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "validation passed",
			fields: fields{
				Email:    "aa@gmail.com",
				Password: "123456",
			},
			assertion: assert.NoError,
		}, {
			name: "validation passed",
			fields: fields{
				Email:    "aa@gmail.com",
				Password: "",
			},
			assertion: assert.NoError,
		}, {
			name: "validation failed",
			fields: fields{
				Email:    "",
				Password: "",
			},
			assertion: assert.Error,
		}, {
			name: "validateion failed",
			fields: fields{
				Email:    "bddfdc.",
				Password: "test123",
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := UserInput{
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
			}
			tt.assertion(t, user.Validate())
		})
	}
}
