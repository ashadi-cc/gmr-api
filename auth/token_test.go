package auth

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

type userTest struct {
	Id    int
	Email string
	Group string
}

func (user userTest) GetUserId() int {
	return user.Id
}

func (user userTest) GetEmail() string {
	return user.Email
}

func (user userTest) GetGroup() string {
	return user.Group
}

func (user *userTest) SetUserId(id int) {
	user.Id = id
}

func (user *userTest) SetEmail(email string) {
	user.Email = email
}

func (user *userTest) SetGroup(group string) {
	user.Group = group
}

func TestCreateToken(t *testing.T) {
	type args struct {
		user UserInterface
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "valid token",
			args: args{
				user: &userTest{Id: 1, Email: "ashadi@mail.com", Group: "user"},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.user)
			tt.assertion(t, err)
			if err == nil {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		user UserInterface
	}
	tests := []struct {
		name      string
		args      args
		want      jwt.MapClaims
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "valid token",
			args: args{
				user: &userTest{Id: 20, Email: "ashadi@gmail", Group: "admin"},
			},
			assertion: assert.NoError,
			want:      jwt.MapClaims{"user_id": 20, "email": "ashadi@gmail", "group": "admin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenString, _ := CreateToken(tt.args.user)
			got, err := ValidateToken(tokenString)
			tt.assertion(t, err)
			assert.Equal(t, fmt.Sprint(got["user_id"]), fmt.Sprint(tt.want["user_id"]))
			assert.Equal(t, got["email"], tt.want["email"])
			assert.Equal(t, got["group"], tt.want["group"])
		})
	}
}

func TestClaimToUser(t *testing.T) {
	type args struct {
		claim jwt.MapClaims
		user  UserInterface
	}
	type want struct {
		id    int
		email string
		group string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid map claim",
			args: args{
				claim: jwt.MapClaims{
					"user_id": float64(2),
					"email":   "ashadi",
					"group":   "user",
				},
				user: &userTest{},
			},
			want: want{
				id:    2,
				email: "ashadi",
				group: "user",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClaimToUser(tt.args.claim, tt.args.user)
			assert.Equal(t, tt.want.id, tt.args.user.GetUserId())
			assert.Equal(t, tt.want.email, tt.args.user.GetEmail())
			assert.Equal(t, tt.want.group, tt.args.user.GetGroup())
		})
	}
}
