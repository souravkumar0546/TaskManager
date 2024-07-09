package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Task struct {
	// gorm.Model
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty" gorm:"not null"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty" gorm:"not null"`
}
