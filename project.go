package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID        uuid.UUID `gorm:"column:project_id;type:char(36)" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"userId"`
	Name      string    `gorm:"column:project_name" json:"name"`
	CreatedAt time.Time `gorm:"column:project_created_at" json:"createdAt"`
	Files     []File    `gorm:"constraint:OnDelete:CASCADE" json:"files"`
	Folders   []Folder  `gorm:"constraint:OnDelete:CASCADE" json:"folder"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if project.ID == uuid.Nil {
		project.ID, _ = uuid.NewV1()
	}
	return
}
