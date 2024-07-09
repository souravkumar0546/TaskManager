package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"task-manager-backend/internal_errors"
	"task-manager-backend/models"
	"task-manager-backend/utils"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID, err := fetch_user_id(r)
	if err != nil {
		writeResponse(w, nil, err)
		return
	}
	rows, err := utils.DB.Query("SELECT id, title, description, status FROM tasks WHERE user_id=$1", userID)
	if err != nil {
		log.Println(err)
		writeResponse(w, nil, internal_errors.ErrInternalError)
		return
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			writeResponse(w, nil, internal_errors.ErrInternalError)
			return
		}
		tasks = append(tasks, task)
	}

	resp := &GetTasksResponse{
		Tasks: tasks,
	}
	writeResponse(w, resp, nil)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, err := fetch_user_id(r)
	if err != nil {
		writeResponse(w, nil, err)
		return
	}
	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload)
		return
	}

	_, err = utils.DB.Exec("INSERT INTO tasks (title, description, status, user_id) VALUES ($1, $2, $3, $4)", task.Title, task.Description, task.Status, userID)
	if err != nil {
		log.Println(err)
		writeResponse(w, nil, internal_errors.ErrInternalError)
		return
	}

	resp := &GetTaskResponse{
		Task: &task,
	}

	writeResponse(w, resp, nil)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid Task ID")
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload)
		return
	}

	userID, err := fetch_user_id(r)
	if err != nil {
		writeResponse(w, nil, err)
		return
	}

	var req UpdateTaskRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload)
		return
	}

	_, err = utils.DB.Exec("UPDATE tasks SET status=$1 WHERE id=$2 AND user_id=$3", req.TaskStatus, id, userID)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInternalError)
		return
	}

	resp := &GetTaskResponse{
		Task: &models.Task{
			ID:     id,
			Status: req.TaskStatus,
		},
	}

	writeResponse(w, resp, nil)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid Task ID")
		writeResponse(w, nil, internal_errors.ErrInvalidRequestPayload)
		return
	}

	userID, err := fetch_user_id(r)
	if err != nil {
		writeResponse(w, nil, err)
		return
	}
	_, err = utils.DB.Exec("DELETE FROM tasks WHERE id=$1 AND user_id=$2", id, userID)
	if err != nil {
		writeResponse(w, nil, internal_errors.ErrInternalError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
