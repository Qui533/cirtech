package enum

import "database/sql/driver"

type PrinterStatus string

const (
	POWEROFF  PrinterStatus = "POWEROFF"
	RUNNING   PrinterStatus = "RUNNING"
	AVAILABLE PrinterStatus = "AVAILABLE"
)

func (s PrinterStatus) String() string {
	switch s {
	case RUNNING:
		return "RUNNING"
	case AVAILABLE:
		return "AVAILABLE"
	}
	return "POWEROFF"
}

func (s *PrinterStatus) Scan(value interface{}) error {
	*s = PrinterStatus(value.([]byte))
	return nil
}

func (s PrinterStatus) Value() (driver.Value, error) {
	return string(s), nil
}
