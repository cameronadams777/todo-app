package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
	CompletedAt time.Time
	UserID      int
}
