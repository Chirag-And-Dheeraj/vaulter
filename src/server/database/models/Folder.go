package models

import (
	"time"

	"gorm.io/gorm"
)

type Folder struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name         string         `json:"name"`
	Hash         string         `json:"hash"`
	Owner        uint           `json:"owner"`
	ParentFolder string         `json:"parent_folder" gorm:"default:'/'"`
}
