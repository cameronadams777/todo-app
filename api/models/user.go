package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"VARCHAR(255); not null"`
	LastName  string `gorm:"VARCHAR(255); not null"`
	Email     string `gorm:"unique; not null"`
	Password  string `gorm:"VARCHAR(255); not null"`
}
