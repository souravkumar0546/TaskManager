package internal_errors

import (
	"errors"
	"net/http"
)

var ErrEmailAlreadyExists = errors.New("email_already_exists")
var ErrInvalidEmail = errors.New("invalid_email")
var ErrInvalidPassword = errors.New("invalid_password")
var ErrInvalidRequestPayload = errors.New("invalid_request_payload")
var ErrInternalError = errors.New("error in internal method")

func GetStatusFromError(err error) int {
	if err == ErrEmailAlreadyExists {
		return http.StatusConflict
	} else if err == ErrInvalidEmail {
		return http.StatusBadRequest
	} else if err == ErrInvalidPassword {
		return http.StatusExpectationFailed
	} else if err == ErrInvalidRequestPayload {
		return http.StatusBadRequest
	} else if err == ErrInternalError {
		return http.StatusInternalServerError
	} else {
		return http.StatusBadRequest
	}
}
