package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserLogin_Validate(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}

	tests := []struct {
		name      string
		fields    fields
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "validation passed",
			fields:    fields{"test", "123"},
			assertion: assert.NoError,
		},
		{
			name:      "validation failed empty",
			fields:    fields{"", ""},
			assertion: assert.Error,
		},
		{
			name:      "validation failed empty username",
			fields:    fields{"", "123"},
			assertion: assert.Error,
		},
		{
			name:      "validation passed empty password",
			fields:    fields{"test", ""},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := UserLogin{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}

			err := user.Validate()
			tt.assertion(t, err)
		})
	}
}
