package models

import (
	"time"

	"gorm.io/gorm"
)

// User gorm model
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}
