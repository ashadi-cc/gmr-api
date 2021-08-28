package util

import (
	"api-gmr/model"
	"encoding/json"
	"log"
	"net/http"
)

type userError struct {
	message string
	err     error
	code    int
}

//Error implementing error interface
func (u userError) Error() string {
	return u.message
}

//NewUserError returns a new userError instance by given values
func NewUserError(code int, message string, err error) *userError {
	if code == 0 {
		code = http.StatusBadRequest
	}

	return &userError{
		message: message,
		err:     err,
		code:    code,
	}
}

//PrintUserError method for write common error to writer
func PrintUserError(w http.ResponseWriter, err error) {
	uError, ok := err.(*userError)
	if ok {
		if uError.err != nil {
			log.Println(uError.err)
		}

		w.WriteHeader(uError.code)
		data := model.CommonMessage{Success: false, Message: uError.message}.WithError(uError.err)
		json.NewEncoder(w).Encode(data)
		return
	}

	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	data := model.CommonMessage{Success: false, Message: "internal server error"}.WithError(err)
	json.NewEncoder(w).Encode(data)
}
