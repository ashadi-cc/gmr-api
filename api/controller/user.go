package controller

import (
	"api-gmr/api/middleware"
	"api-gmr/model"
	"api-gmr/service"
	"api-gmr/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

//User represents User Controller
type User struct {
	userService service.IUserService
	maxMemory   int64
}

//NewUser returns new User instance
func NewUser(userService service.IUserService) *User {
	return &User{
		userService: userService,
		maxMemory:   1024,
	}
}

//Info User info handler method
func (u User) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		util.PrintUserError(w, fmt.Errorf("can't load user from context"))
		return
	}

	userInfo, err := u.userService.UserInfo(userCtx.Id)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: userInfo})

}

//Update User update hander method
func (u User) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		util.PrintUserError(w, fmt.Errorf("can't load user from context"))
		return
	}

	var userInput model.UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		userError := util.NewUserError(http.StatusBadRequest, "invalid payload", err)
		util.PrintUserError(w, userError)
		return
	}

	if err := userInput.Validate(); err != nil {
		userError := util.NewUserError(http.StatusBadRequest, "invalid payload", err)
		util.PrintUserError(w, userError)
		return
	}

	userCtx.Email = userInput.Email
	userCtx.Password = userInput.Password

	if err = u.userService.UpdateUser(userCtx); err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Message: "user updated"})
}

//Billing user billing handler
func (u User) Billing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		util.PrintUserError(w, fmt.Errorf("can't load user from context"))
		return
	}

	data, err := u.userService.GetBilling(userCtx)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: data})
}

//Upload upload image user handler
func (u User) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	// if !ok {
	// 	util.PrintUserError(w, fmt.Errorf("can't load user from context"))
	// 	return
	// }

	r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024) // 2 Mb
	if err := r.ParseMultipartForm(u.maxMemory); err != nil {
		util.PrintUserError(w, err)
		return
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		util.PrintUserError(w, err)
		return
	}
	defer uploadedFile.Close()

	mType, err := mimetype.DetectReader(uploadedFile)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	if !strings.HasPrefix(mType.String(), "image") {
		util.PrintUserError(w, util.NewUserError(http.StatusBadRequest, "file must an image", nil))
		return
	}

	filename := handler.Filename //you should change the name
	//fileExt := filepath.Ext(handler.Filename)

	dir, _ := os.Getwd()
	fileLocation := filepath.Join(dir, "data", "upload", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		util.PrintUserError(w, err)
		return
	}

	// description := r.FormValue("description")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true})
}
