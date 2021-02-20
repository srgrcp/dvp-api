package models

import (
	"time"

	"github.com/srgrcp/dvp-api/modules/user/models"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

// Ticket gorm model
type Ticket struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Status string `json:"status"`
	UserID uint   `json:"-"`

	User models.User `json:"user"`
}

// CreateTicketInput create ticket endpoint parameters
type CreateTicketInput struct {
	Status *string `json:"status" validate:"oneof=abierto cerrado"`
	UserID uint    `json:"userId"`
}

// IsValid input validation
func (t *CreateTicketInput) IsValid() bool {
	if t.Status == nil {
		status := "abierto"
		t.Status = &status
	}
	validator := validator.New()
	err := validator.Struct(t)
	return err == nil
}

// UpdateTicketInput update ticket endpoint parameters
type UpdateTicketInput struct {
	Status *string `json:"status" validate:"oneof=abierto cerrado"`
	UserID *uint   `json:"userId"`
}

// IsValid input validation
func (t *UpdateTicketInput) IsValid() bool {
	validator := validator.New()
	err := validator.Struct(t)
	isValid := err == nil
	return t.Status == nil || isValid
}
