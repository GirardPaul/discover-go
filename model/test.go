package test

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model            // Adds some metadata fields to the table
	ID          uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title       string    `gorm:"not null;type:varchar(100);default: null"`
	Description string    `gorm:"not null;type:varchar(100);default: null"`
}
