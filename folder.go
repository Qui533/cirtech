package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	ID        uuid.UUID  `gorm:"type:char(36);column:folder_id" json:"id"`
	ParentID  *uuid.UUID `gorm:"type:char(36)" json:"parentId"`
	ProjectID *uuid.UUID `gorm:"type:char(36)" json:"projectId"`
	Name      string     `gorm:"column:folder_name" json:"name"`
	CreatedAt time.Time  `gorm:"column:folder_created_at" json:"createdAt"`
	Files     []File     `gorm:"constraint:OnDelete:CASCADE" json:"files"`
	Children  []Folder   `gorm:"foreignkey:ParentID;constraint:OnDelete:CASCADE" json:"children"`
}

func (folder *Folder) BeforeCreate(tx *gorm.DB) (err error) {
	if folder.ID == uuid.Nil {
		folder.ID, _ = uuid.NewV1()
	}
	return
}
