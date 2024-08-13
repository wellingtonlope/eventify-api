package user

import (
	"time"

	"github.com/wellingtonlope/eventify-api/internal/business/tag"
)

type (
	User struct {
		ID    string
		Name  string
		Rule  string
		Email string
		Tags  []tag.Tag

		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
