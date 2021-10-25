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

//User represent User mock repository
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

func (user *User) assertionMethodCall(method string, expectedCalls int, arguments ...interface{}) {
	if expectedCalls > 0 {
		user.AssertCalled(user.t, method, arguments...)
	}
	user.AssertNumberOfCalls(user.t, method, expectedCalls)
}

//ListenOnFindByUsername listening FindByUsername call and return values from args
func (user *User) ListenOnFindByUsername(returnUser repository.UserModel, returnErr error) *User {
	user.On(FindByUsernameMethod, mock.Anything, mock.IsType(string(""))).Return(returnUser, returnErr)
	return user
}

//AssertFindByUsernameCall assertion FindByUsername call
func (user *User) AssertFindByUsernameCall(expectedCalls int, arguments ...interface{}) *User {
	user.assertionMethodCall(FindByUsernameMethod, expectedCalls, arguments...)
	return user
}

//ListenOnFindByUserID listening FindByUserID call and returns values from args
func (user *User) ListenOnFindByUserID(returnUser repository.UserModel, returnErr error) *User {
	user.On(FindByUserIDMethod, mock.Anything, mock.IsType(int(0))).Return(returnUser, returnErr)
	return user
}

//AssertFindByUserIDCall assertion FindByUserID call
func (user *User) AssertFindByUserIDCall(expectedCalls int, arguments ...interface{}) *User {
	user.assertionMethodCall(FindByUserIDMethod, expectedCalls, arguments...)
	return user
}

//ListenOnUpdateEmailandPassword listening on UpdateEmailandPassword call and returns values from args
func (user *User) ListenOnUpdateEmailandPassword(returnErr error) *User {
	user.On(UpdateEmailandPasswordMethod, mock.Anything, mock.Anything).Return(returnErr)
	return user
}

//AssertUpdateEmailandPasswordCall assertion UpdateEmailandPassword call
func (user *User) AssertUpdateEmailandPasswordCall(expectedCalls int, arguments ...interface{}) *User {
	user.assertionMethodCall(UpdateEmailandPasswordMethod, expectedCalls, arguments...)
	return user
}
