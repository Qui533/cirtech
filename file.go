package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type File struct {
	ID          uuid.UUID     `gorm:"type:char(36);column:file_id" json:"id"`
	FolderID    uuid.NullUUID `gorm:"type:char(36)" json:"folderId"`
	ProjectID   uuid.NullUUID `gorm:"type:char(36)" json:"projectId"`
	Name        string        `gorm:"column:file_name" json:"name"`
	Extension   string        `gorm:"column:file_extension" json:"extension"`
	Size        int           `gorm:"column:file_size" json:"size"`
	URL         string        `gorm:"column:file_url" json:"url"`
	Application string        `gorm:"column:file_application" json:"application"`
	CreatedAt   time.Time     `gorm:"column:file_created_at" json:"createdAt"`
}

func (file *File) BeforeCreate(tx *gorm.DB) (err error) {
	if file.ID == uuid.Nil {
		file.ID, _ = uuid.NewV1()
	}
	return
}
