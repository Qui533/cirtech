package model

import (
	"3d-print-inventory/enum"

	"github.com/gofrs/uuid"
)

type Permission struct {
	PrinterID uuid.UUID       `gorm:"primaryKey;column:printer_id" json:"printerId"`
	UserID    uuid.UUID       `gorm:"primaryKey;column:user_id" json:"userId"`
	Name      enum.Permission `gorm:"type:enum('OWN', 'EDIT', 'VIEW');column:permission_name" json:"name"`
}
