package repository

import "github.com/dnevsky/restaurant-back/internal/models"

type UserRepo interface {
	Find(id uint) (user models.User, err error)
	Create(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
}
