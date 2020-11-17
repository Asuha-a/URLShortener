package db

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

// URL Model
type URL struct {
	UUID       uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	UserID     uuid.UUID
	Title      string
	URL        string `gorm:"unique"`
	RedirectTo datatypes.JSON
	CreatedAt  time.Time
}
