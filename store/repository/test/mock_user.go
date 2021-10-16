package test

import (
	"api-gmr/store/repository"
	"api-gmr/store/repository/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

const (
	//FindByUsernameMethod represent FindByUsername name method
	FindByUsernameMethod = "FindByUsername"
	//FindByUserIDMethod represent FindByUserID name method
	FindByUserIDMethod = "FindByUserID"
	//UpdateEmailandPasswordMethod represent UpdateEmailandPassword name method
	UpdateEmailandPasswordMethod = "UpdateEmailandPassword"
)

type User struct {
	t *testing.T
	*mocks.User
}

//NewMockUser returns new instance user mock
func NewMockUser(t *testing.T) *User {
	return &User{
		t:    t,
		User: &mocks.User{},
	}
}

//ListenOnFindByUsername listening FindByUsername call and return values from args
func (user *User) ListenOnFindByUsername(returnUser repository.UserModel, returnErr error) *User {
	user.On(FindByUsernameMethod, mock.Anything, mock.IsType(string(""))).Return(returnUser, returnErr)
	return user
}

//AssertFindByUsernameCall assertion FindByUsername call
func (user *User) AssertFindByUsernameCall(expectedCalls int, arguments ...interface{}) *User {
	if expectedCalls > 0 {
		user.AssertCalled(user.t, FindByUsernameMethod, arguments...)
	}
	user.AssertNumberOfCalls(user.t, FindByUsernameMethod, expectedCalls)
	return user
}
