package repository

import (
	categoryDTO "github.com/dnevsky/restaurant-back/internal/dto/category"
	"github.com/dnevsky/restaurant-back/internal/models"
)

type CategoryRepo interface {
	Create(category models.Category) (models.Category, error)
	Update(category *models.Category) error
	Find(id uint) (category models.Category, err error)
	Delete(id uint) error
	List(dto categoryDTO.CategoryListDTO) (categories []models.Category, err error)
}
