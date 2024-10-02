package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
// func writeErrorResponse(w http.ResponseWriter, err error) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("APP_URL"))
// 	statusCode := internal_errors.GetStatusFromError(err)
// 	w.WriteHeader(statusCode)
// 	type genericError struct {
// 		Message string `json:"message"`
// 	}
// 	errResponse := genericError{
// 		Message: err.Error(),
// 	}
// 	if writeErr := json.NewEncoder(w).Encode(errResponse); writeErr != nil {
// 		zap.NewExample().Error("could not write response", zap.Error(writeErr))
// 	}
// }
func writeErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	// Set allowed origins based on your environment
	appURL := os.Getenv("APP_URL")
	localhost := "http://localhost:3000"
	
	// You can dynamically decide if you're in development or production
	allowedOrigins := fmt.Sprintf("%s, %s", appURL, localhost)
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)

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
