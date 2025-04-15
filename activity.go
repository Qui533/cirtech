package model

import (
	"3d-print-inventory/enum"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	ID                   uuid.UUID           `gorm:"type:char(36);column:activity_id" json:"id"`
	PrinterActivationKey string              `gorm:"type:char(17);column:printer_activation_key" json:"printerActivationKey"`
	EstimatedTime        int                 `gorm:"column:activity_estimated_time" json:"estimatedTime"`
	FileID               uuid.UUID           `gorm:"column:file_id;type:char(36);" json:"fileId"`
	File                 File                `gorm:"constraint:OnDelete:CASCADE" json:"file"`
	Status               enum.ActivityStatus `gorm:"column:activity_status;type:enum('PENDING', 'PRINTING', 'DONE');default:'PENDING'" json:"status"`
	CreatedAt            time.Time           `gorm:"column:activity_created_at" json:"created_at"`
}

func (activity *Activity) BeforeCreate(tx *gorm.DB) (err error) {
	if activity.ID == uuid.Nil {
		activity.ID, _ = uuid.NewV1()
	}
	return
}
