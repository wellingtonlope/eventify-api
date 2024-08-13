package ticket

import "time"

type (
	Status string
	Ticket struct {
		ID      string
		EventID string
		Status  string
		OwnerID string

		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

const (
	StatusRequested  = Status("requested")
	StatusApproved   = Status("approved")
	StatusUtilized   = Status("utilized")
	StatusUnutilized = Status("unutilized")
	StatusRejected   = Status("rejected")
	StatusCancelled  = Status("cancelled")
)
