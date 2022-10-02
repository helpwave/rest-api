package models

import (
	"github.com/google/uuid"
	"time"
)

type Emergency struct {
	ID        uuid.UUID
	StartLoc  Point
	TimeStamp time.Time
}
