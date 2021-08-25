package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHash(password string) string {
	hash, _ := HashPassword(password)
	return hash
}

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "valid password",
			args: args{"computer"},

			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			tt.assertion(t, err)
			if err != nil {
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid hash",
			args: args{"computer", testHash("computer")},
			want: true,
		}, {
			name: "valid laravel hash",
			args: args{"computer", "$2y$10$mzLnmcoii8PwUhhqX8NzDuLShJ0F67woPFwEkU1XfvOPBCal2FnYK"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CheckPasswordHash(tt.args.password, tt.args.hash))
		})
	}
}
