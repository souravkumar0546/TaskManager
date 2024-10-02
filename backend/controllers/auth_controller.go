package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"

	"task-manager-backend/internal_errors"
	"task-manager-backend/models"
	"task-manager-backend/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("my_secret_key")
var defaultAvatar = os.Getenv("API_URL") + "/avatars/defaultpic.png"

type Claims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload, r)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password")
		writeResponse(w, nil, internal_errors.ErrInternalError, r)
		return
	}

	avatar := defaultAvatar
	// Corrected SQL statement to include the avatar column
	_, err = utils.DB.Exec("INSERT INTO users (email, name, password, avatar) VALUES ($1, $2, $3, $4)", req.Email, req.Name, hashedPassword, avatar)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				writeResponse(w, nil, internal_errors.ErrEmailAlreadyExists, r)
				return
			}
		}
		log.Println(err)
		writeResponse(w, nil, internal_errors.ErrInternalError, r)
		return
	}

	resp := &SignUpResponse{
		Email: req.Email,
	}
	writeResponse(w, resp, nil, r)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LogInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload, r)
		return
	}

	var storedCreds models.User
	err = utils.DB.QueryRow("SELECT id, email, name, password FROM users WHERE email=$1", req.Email).Scan(&storedCreds.ID, &storedCreds.Email, &storedCreds.Name, &storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			writeResponse(w, nil, internal_errors.ErrInvalidEmail, r)
			return
		}
		writeResponse(w, nil, internal_errors.ErrInternalError, r)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(req.Password)); err != nil {
		writeResponse(w, nil, internal_errors.ErrInvalidPassword, r)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: storedCreds.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInternalError, r)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,                  // Ensure this is true for HTTPS
		SameSite: http.SameSiteNoneMode, // SameSite: None for cross-site requests
	})

	storedCreds.Password = ""
	resp := &LogInResponse{
		User: &storedCreds,
	}
	writeResponse(w, resp, nil, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Set the cookie with an expired time to clear it
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Set an expired time
		HttpOnly: true,
		Secure:   true,                  // Ensure this is true for HTTPS
		SameSite: http.SameSiteNoneMode, // SameSite: None for cross-site requests
	})

	// Send a response indicating successful logout
	w.WriteHeader(http.StatusOK)
	writeResponse(w, map[string]string{"message": "Logged out successfully"}, nil, r)
}