package repository

import (
	foodDTO "github.com/dnevsky/restaurant-back/internal/dto/food"
	"github.com/dnevsky/restaurant-back/internal/models"
)

type FoodRepo interface {
	Create(food models.Food) (models.Food, error)
	Update(food *models.Food) error
	Find(id uint) (food models.Food, err error)
	Delete(id uint) error
	List(dto foodDTO.FoodListDTO) (foods []models.Food, err error)
}
