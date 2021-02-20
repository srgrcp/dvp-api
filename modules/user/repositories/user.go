package repositories

import (
	"errors"
	"fmt"

	"github.com/srgrcp/dvp-api/modules/user/models"
	"github.com/srgrcp/dvp-api/utils"
	"gorm.io/gorm"
)

// UserRepository db handler
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns new TicketRepository instance
func NewUserRepository() *UserRepository {
	db := utils.MySQLConnection()
	return &UserRepository{db: db}
}

// Find a user by id
func (r *UserRepository) Find(userID uint) (*models.User, error) {
	user := &models.User{}
	err := r.db.First(user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user does not exists")
		}
		return nil, err
	}
	return user, nil
}
