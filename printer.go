package model

import (
	"3d-print-inventory/enum"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Printer struct {
	ID                   uuid.UUID          `gorm:"type:char(36);column:printer_id" json:"id"`
	ActivationKey        string             `gorm:"type:char(17);column:printer_activation_key" json:"activationKey"`
	Name                 string             `gorm:"column:printer_name" json:"name"`
	UserID               uuid.NullUUID      `gorm:"type:char(36)" json:"userId"`
	Status               enum.PrinterStatus `gorm:"column:printer_status;type:enum('POWEROFF', 'AVAILABLE', 'RUNNING');default:'AVAILABLE'" json:"status"`
	Model                string             `gorm:"column:printer_model" json:"model"`
	Technology           string             `gorm:"column:printer_technology" json:"technology"`
	ConsumablesDiameter  string             `gorm:"column:printer_consumables_diameter" json:"consumablesDiameter"`
	SliceThickness       string             `gorm:"column:printer_slice_thickness" json:"slice_thickness"`
	MaximumPrintingSpeed string             `gorm:"column:printer_maximum_printing_speed" json:"maximumPrintingSpeed"`
	Filament             string             `gorm:"column:printer_filament" json:"filament"`
	NozzleDiameter       string             `gorm:"column:printer_nozzle_diameter" json:"nozzleDiameter"`
	MaximumPrintSize     string             `gorm:"column:printer_maximum_print_size" json:"maximumPrintSize"`
}

func (printer *Printer) BeforeCreate(tx *gorm.DB) (err error) {
	if printer.ID == uuid.Nil {
		printer.ID, _ = uuid.NewV1()
	}
	return
}
