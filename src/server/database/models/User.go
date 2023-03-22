package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FirstName    string         `json:"first_name" gorm:"not null;default:null"`
	LastName     string         `json:"last_name" `
	Email        string         `json:"email" gorm:"unique;not null;default:null"`
	Pin          string         `json:"pin" gorm:"not null;default:null"`
	ProfilePhoto string         `json:"profile_photo"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
}
