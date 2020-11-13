package db

import (
	uuid "github.com/satori/go.uuid"
)

// User Model
type User struct {
	UUID       uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	EMAIL      string    `gorm:"unique"`
	PERMISSION string    `gorm:"->"`
	PASSWORD   string
}
