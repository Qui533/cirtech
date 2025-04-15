package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Reconstruction struct {
	ID        uuid.UUID `gorm:"type:char(36);column:reconstruction_id" json:"id"`
	ImageID   uuid.UUID `gorm:"type:char(36);column:reconstruction_image_id" json:"imageId"`
	ExtractID uuid.UUID `gorm:"type:char(36);column:reconstruction_extract_id" json:"extractId"`
	ResultID  uuid.UUID `gorm:"type:char(36);column:reconstruction_result_id" json:"resultId"`
}

func (reconstruction *Reconstruction) BeforeCreate(tx *gorm.DB) (err error) {
	if reconstruction.ID == uuid.Nil {
		reconstruction.ID, _ = uuid.NewV1()
	}
	return
}
