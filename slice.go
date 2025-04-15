package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Slice struct {
	ID       uuid.UUID `gorm:"type:char(36);column:slice_id" json:"id"`
	FileID   uuid.UUID `gorm:"type:char(36)" json:"fileId"`
	ResultID uuid.UUID `gorm:"type:char(36)" json:"resultId"`
}

func (slice *Slice) BeforeCreate(tx *gorm.DB) (err error) {
	if slice.ID == uuid.Nil {
		slice.ID, _ = uuid.NewV1()
	}
	return
}
