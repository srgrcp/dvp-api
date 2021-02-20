package repositories

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/morkid/paginate"
	"github.com/srgrcp/dvp-api/modules/ticket/models"
	"github.com/srgrcp/dvp-api/modules/user/repositories"
	"github.com/srgrcp/dvp-api/utils"
	"gorm.io/gorm"
)

// TicketRepository db handler
type TicketRepository struct {
	db *gorm.DB
}

// NewTicketRepository returns new TicketRepository instance
func NewTicketRepository() *TicketRepository {
	db := utils.MySQLConnection()
	return &TicketRepository{db: db}
}

// Create a new Ticket in database
func (r *TicketRepository) Create(ticketInput *models.CreateTicketInput) (*models.Ticket, error) {
	if !ticketInput.IsValid() {
		return nil, fmt.Errorf("invalid status value, only acepted: 'abierto' or 'cerrado'")
	}
	ticket := &models.Ticket{
		Status: *ticketInput.Status,
		UserID: ticketInput.UserID,
	}
	err := r.db.Save(ticket).Error
	if err != nil {
		return nil, err
	}
	ticket, err = r.Find(fmt.Sprint(ticket.ID))
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

// Query paginated and filtered list of tickets
func (r *TicketRepository) Query(req *http.Request) paginate.Page {
	page := paginate.New().Response(r.db.Preload("User").Model(&models.Ticket{}), req, &[]models.Ticket{})
	return page
}

// Update information of an existing ticket
func (r *TicketRepository) Update(ticketID string, ticketInput *models.UpdateTicketInput) (*models.Ticket, error) {
	if !ticketInput.IsValid() {
		return nil, fmt.Errorf("invalid status value, only acepted: 'abierto' or 'cerrado'")
	}

	ticket, err := r.Find(ticketID)
	if err != nil {
		return nil, err
	}

	if ticketInput.Status != nil {
		ticket.Status = *ticketInput.Status
	}
	if ticketInput.UserID != nil {
		ticket.UserID = *ticketInput.UserID
		user, err := repositories.NewUserRepository().Find(ticket.UserID)
		if err != nil {
			return nil, err
		}
		ticket.User = *user
	}

	err = r.db.Save(ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

// Find a ticket by id
func (r *TicketRepository) Find(ticketID string) (*models.Ticket, error) {
	ticket := &models.Ticket{}
	err := r.db.Preload("User").First(ticket, ticketID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ticket does not exists")
		}
		return nil, err
	}
	return ticket, nil
}

// Delete ticket by id
func (r *TicketRepository) Delete(ticketID string) (*models.Ticket, error) {
	ticket, err := r.Find(ticketID)
	if err != nil {
		return nil, err
	}
	if err = r.db.Delete(ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}
