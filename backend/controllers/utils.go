package controllers

import (
	"encoding/json"
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

// writeResponse receives the response and error and appropriately fills the response status, header, and body
func writeResponse(w http.ResponseWriter, response interface{}, err error, r *http.Request) { // Added r *http.Request here
	if err != nil {
		writeErrorResponse(w, r, err) // Pass the request here
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Check if response is not nil before encoding
	if response != nil {
		if writeErr := json.NewEncoder(w).Encode(response); writeErr != nil {
			zap.NewExample().Error("could not write response", zap.Error(writeErr))
		}
	}
}

// writeErrorResponse creates an error response and writes to the user.
func writeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	// Get the origin of the request
	origin := r.Header.Get("Origin")
	
	// Allow specific origins
	appURL := os.Getenv("APP_URL") // Netlify URL
	localhost := "http://localhost:3000"

	// Check if the origin is allowed
	if origin == appURL || origin == localhost {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

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