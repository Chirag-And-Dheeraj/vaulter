package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID               uint           `json:"id" gorm:"primarykey"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name             string         `json:"name"`
	Hash             string         `json:"hash"`
	Size             int            `json:"size"`
	Mime             string         `json:"mime"`
	Owner            uint           `json:"owner"`
	IsFolder         bool           `json:"is_folder"`
	IsPublic         bool           `json:"is_public"`
	IsUploadComplete bool           `json:"is_upload_complete" gorm:"default:false"`
	ParentFolder     string         `json:"parent_folder" gorm:"default:'/'"`
}
