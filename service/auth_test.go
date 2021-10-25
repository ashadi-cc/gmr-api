package service

import (
	"api-gmr/model"
	"api-gmr/store/repository"
	"api-gmr/store/repository/test"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Validate(t *testing.T) {
	type args struct {
		data      model.UserLogin
		returnErr error
		user      repository.UserModel
	}
	tests := []struct {
		name                   string
		args                   args
		want                   model.User
		numberFinduserNameCall int
		assertion              assert.ErrorAssertionFunc
	}{
		{
			name: "valid password",
			args: args{
				data: model.UserLogin{
					Username: "admin",
					Password: "computer",
				},
				returnErr: nil,
				user: model.User{
					Id:       1,
					Username: "admin",
					Password: "$2y$10$mzLnmcoii8PwUhhqX8NzDuLShJ0F67woPFwEkU1XfvOPBCal2FnYK",
				},
			},
			want: model.User{
				Id: 1,
			},
			numberFinduserNameCall: 1,
			assertion:              assert.NoError,
		},
		{
			name: "invalid password",
			args: args{
				data: model.UserLogin{
					Username: "admin",
					Password: "password",
				},
				returnErr: nil,
				user: model.User{
					Id:       1,
					Username: "admin",
					Password: "invalid",
				},
			},
			want: model.User{
				Id: 0,
			},
			numberFinduserNameCall: 1,
			assertion:              assert.Error,
		},
		{
			name: "invalid username",
			args: args{
				data: model.UserLogin{
					Username: "invalidusername",
					Password: "password",
				},
				returnErr: fmt.Errorf("user not found"),
				user:      nil,
			},
			want:                   model.User{Id: 0},
			numberFinduserNameCall: 1,
			assertion:              assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := test.NewMockUser(t)
			repo.ListenOnFindByUsername(tt.args.user, tt.args.returnErr)

			service := &AuthService{
				userRepo: repo,
			}
			got, err := service.Validate(tt.args.data)

			repo.AssertExpectations(t)
			repo.AssertFindByUsernameCall(tt.numberFinduserNameCall, mock.Anything, tt.args.data.Username)
			tt.assertion(t, err)
			assert.Equal(t, tt.want.Id, got.Id)
		})
	}
}
