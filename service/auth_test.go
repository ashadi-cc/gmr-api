package service

import (
	"api-gmr/model"
	"api-gmr/store/repository/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Validate(t *testing.T) {
	type args struct {
		data      model.UserLogin
		returnErr error
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
			},
			want: model.User{
				Id:       1,
				Username: "admin",
				Password: "$2y$10$mzLnmcoii8PwUhhqX8NzDuLShJ0F67woPFwEkU1XfvOPBCal2FnYK",
			},
			numberFinduserNameCall: 1,
			assertion:              assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := test.NewMockUser(t)
			repo.ListenOnFindByUsername(tt.want, tt.args.returnErr)

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
