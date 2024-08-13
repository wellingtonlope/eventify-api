package event

import (
	"time"

	"github.com/wellingtonlope/eventify-api/internal/business/tag"
)

type (
	Status string
	Event  struct {
		ID             string
		Name           string
		Description    string
		Status         string
		Capacity       int
		Tags           []tag.Tag
		EventManagerID string

		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

const (
	StatusCreated   = Status("created")
	StatusOpened    = Status("opened")
	StatusFinished  = Status("finished")
	StatusCancelled = Status("cancelled")
)
