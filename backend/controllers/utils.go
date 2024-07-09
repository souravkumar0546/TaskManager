package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"task-manager-backend/internal_errors"

	"go.uber.org/zap"
)

func fetch_user_id(r *http.Request) (int, error) {
	userID := r.Context().Value("userID")
	if userID == nil {
		log.Println("User ID missing from context")
		return -1, internal_errors.ErrInvalidRequestPayload
	}
	return userID.(int), nil
}

// writeResponse receives the response and error and appropriately fills the response status, header and body
func writeResponse(w http.ResponseWriter, response interface{}, err error) {
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if writeErr := json.NewEncoder(w).Encode(response); writeErr != nil {
		zap.NewExample().Error("could not write response", zap.Error(writeErr))
	}
}

// writeErrorResponse creates an error response and writes to the user.
func writeErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://tasktrackhub.netlify.app")
	statusCode := internal_errors.GetStatusFromError(err)
	w.WriteHeader(statusCode)
	type genericError struct {
		Message string `json:"message"`
	}
	errResponse := genericError{
		Message: err.Error(),
	}
	if writeErr := json.NewEncoder(w).Encode(errResponse); writeErr != nil {
		zap.NewExample().Error("could not write response", zap.Error(writeErr))
	}
}
