package entity

import (
	"time"

	"github.com/google/uuid"
)

type ActivityLog struct {
	ID           uuid.UUID
	ActivityType string
	CreatedAt    time.Time
}
