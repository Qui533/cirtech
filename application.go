package enum

type Application string

const (
	INVENTORY      Application = "INVENTORY"
	RECONSTRUCTION Application = "RECONSTRUCTION"
)

func (p Application) String() string {
	switch p {
	case INVENTORY:
		return "INVENTORY"
	case RECONSTRUCTION:
		return "RECONSTRUCTION"
	}
	return "UNKNOWN"
}
