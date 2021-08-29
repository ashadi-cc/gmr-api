package service

import (
	"api-gmr/auth"
	"api-gmr/config"
	"api-gmr/model"
	"api-gmr/store/repository"
	"api-gmr/util"
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

//IUserService represents a service for user methods
type IUserService interface {

	//UserInfo returns model.User by given userID
	UserInfo(userID int) (model.User, error)

	//UpdateUser method for update users by given user model payload.
	UpdateUser(user model.User) error

	//GetBilling returns billings for particular user
	GetBilling(user model.User) (model.BillingInfo, error)
}

//UserService impelmenting IUserService
type UserService struct {
	userRepo    repository.User
	billRepo    repository.Billing
	paymentRepo repository.Payment
}

//NewUserService return a new UserService instance
func NewUserService() IUserService {
	return &UserService{
		userRepo:    repo().GetUserRepository(),
		billRepo:    repo().GetBillingRepository(),
		paymentRepo: repo().GetPaymentRepository(),
	}
}

//UserInfo impelemnting IUserService.UserINfo
func (service *UserService) UserInfo(userID int) (model.User, error) {
	var user model.User

	dbUser, err := service.userRepo.FindByUserID(context.Background(), userID)
	if err != nil {
		if cause := errors.Cause(err); cause == sql.ErrNoRows {
			return user, util.NewUserError(http.StatusBadRequest, "user id not found", err)
		}
		return user, err
	}

	user = model.User{
		Id:       dbUser.GetUserID(),
		Email:    dbUser.GetEmail(),
		Group:    dbUser.GetGroup(),
		Username: dbUser.GetUsername(),
		Blok:     dbUser.GetBlok(),
		Name:     dbUser.GetName(),
	}
	return user, nil
}

//UpdateUser implementing IUserService.UpdateUser
func (service *UserService) UpdateUser(user model.User) error {
	if user.Password != "" {
		hashPasword, err := auth.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = hashPasword
	}

	err := service.userRepo.UpdateEmailandPassword(context.Background(), user)
	return err
}

//GetBilling implementing IUserService.GetBilling
func (service *UserService) GetBilling(user model.User) (model.BillingInfo, error) {
	var bInfo model.BillingInfo

	localTime, err := util.TimeIn(time.Now(), config.GetApp().TimeZone)
	if err != nil {
		return bInfo, errors.Wrap(err, "unable to load local timezone")
	}

	billingFilter := model.BillingFilter{
		Year:   localTime.Year(),
		Month:  int(localTime.Month()),
		UserID: user.GetUserID(),
		Status: "B",
	}
	thisMonth, err := service.billRepo.GetBillWithFilter(context.Background(), billingFilter)
	if err != nil {
		return bInfo, err
	}

	otherBill, err := service.billRepo.GetOtherBillWithFilter(context.Background(),
		user.GetUserID(),
		localTime.Year(),
		int(localTime.Month()))
	if err != nil {
		return bInfo, err
	}

	payments, err := service.paymentRepo.All(context.Background())
	if err != nil {
		return bInfo, err
	}

	thisMonthBill := model.BillRepoToBilling(thisMonth)
	otherMonthBill := model.BillRepoToBilling(otherBill)
	paymentsList := model.PaymentRepoToPayments(payments)

	bInfo = model.BillingInfo{
		ThisMonth:     model.ItemBilling{Data: thisMonthBill.Display(), Total: thisMonthBill.TotalAmount()},
		OtherBill:     model.ItemBilling{Data: otherMonthBill.Display(), Total: otherMonthBill.TotalAmount()},
		PaymentMethod: paymentsList,
	}

	return bInfo, nil
}
