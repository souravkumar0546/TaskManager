package controllers

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"task-manager-backend/internal_errors"
	"task-manager-backend/models"
	"task-manager-backend/utils"
)

const (
	uploadsDir = "./user_data/avatar" // Directory where uploaded images will be stored
)

func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	userID, err := fetch_user_id(r)
	if err != nil {
		writeResponse(w, nil, err, r) // Include r as the last parameter
		return
	}

	// Parse multipart form data to get the file
	err = r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload, r) // Include r as the last parameter
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		log.Println("Error retrieving file from form data:", err)
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload, r) // Include r as the last parameter
		return
	}
	defer file.Close()

	// Validate User ID
	err = utils.DB.QueryRow("SELECT id FROM users WHERE id=$1", userID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeResponse(w, nil, internal_errors.ErrInvalidEmail, r) // Include r as the last parameter
			return
		}
		writeResponse(w, nil, internal_errors.ErrInternalError, r) // Include r as the last parameter
		return
	}

	// Generate a unique filename for the uploaded file
	filename := fmt.Sprintf("%d%s", userID, filepath.Ext(handler.Filename))
	filePath := filepath.Join(uploadsDir, filename)

	// Create a new file in the uploads directory
	out, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating file:", err)
		writeResponse(w, nil, internal_errors.ErrInternalError, r) // Include r as the last parameter
		return
	}
	defer out.Close()

	// Copy the file to the destination path
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Error saving file:", err)
		writeResponse(w, nil, internal_errors.ErrInternalError, r) // Include r as the last parameter
		return
	}

	// Update user's avatar URL in the database
	avatarURL := os.Getenv("API_URL") + "/avatars/" + filename

	_, err = utils.DB.Exec("UPDATE users SET avatar=$1 WHERE id=$2", avatarURL, userID)
	if err != nil {
		log.Println("Error updating avatar URL in database:", err)
		writeResponse(w, nil, internal_errors.ErrInternalError, r) // Include r as the last parameter
		return
	}

	// Return the updated avatar URL in the response
	resp := map[string]string{"avatar": avatarURL}
	writeResponse(w, resp, nil, r) // Include r as the last parameter
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := fetch_user_id(r)
	if err != nil {
		writeResponse(w, nil, err, r) // Include r as the last parameter
		return
	}
	var userProfile models.User
	err = utils.DB.QueryRow("SELECT id, email, name, avatar FROM users WHERE id=$1", userID).Scan(&userProfile.ID, &userProfile.Email, &userProfile.Name, &userProfile.Avatar)
	if err != nil {
		if err == sql.ErrNoRows {
			writeResponse(w, nil, internal_errors.ErrInvalidEmail, r) // Include r as the last parameter
			return
		}
		writeResponse(w, nil, internal_errors.ErrInternalError, r) // Include r as the last parameter
		return
	}

	resp := &GetUserProfileResponse{
		User: &userProfile,
	}
	writeResponse(w, resp, nil, r) // Include r as the last parameter
}