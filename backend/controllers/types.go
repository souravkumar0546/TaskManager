package controllers

import "task-manager-backend/models"

type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Email string `json:"email"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInResponse struct {
	User *models.User `json:"user"`
}

type GetUserProfileResponse struct {
	User *models.User `json:"user"`
}

type GetTasksResponse struct {
	Tasks []models.Task `json:"tasks"`
}

type GetTaskResponse struct {
	Task *models.Task `json:"task"`
}

type UpdateTaskRequest struct {
	TaskStatus string `json:"status"`
}
