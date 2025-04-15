package enum

import "database/sql/driver"

type ActivityStatus string

const (
	PENDING  ActivityStatus = "PENDING"
	PRINTING ActivityStatus = "PRINTING"
	DONE     ActivityStatus = "DONE"
)

func (s ActivityStatus) String() string {
	switch s {
	case PRINTING:
		return "PRINTING"
	case DONE:
		return "DONE"
	}
	return "PENDING"
}

func (s *ActivityStatus) Scan(value interface{}) error {
	*s = ActivityStatus(value.([]byte))
	return nil
}

func (s ActivityStatus) Value() (driver.Value, error) {
	return string(s), nil
}
