package postgres

import (
	foodDTO "github.com/dnevsky/restaurant-back/internal/dto/food"
	"github.com/dnevsky/restaurant-back/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FoodRepo struct {
	DB *gorm.DB
}

func NewFoodRepo(db *gorm.DB) *FoodRepo {
	return &FoodRepo{DB: db}
}

func (r *FoodRepo) Create(food models.Food) (models.Food, error) {
	err := r.DB.Clauses(clause.Returning{}).Create(&food).Error
	return food, err
}

func (r *FoodRepo) Update(food *models.Food) error {
	return r.DB.Omit("Category").Save(food).Error
}

func (r *FoodRepo) Find(id uint) (food models.Food, err error) {
	err = r.DB.Preload("Category").First(&food, "id = ?", id).Error
	return food, err
}

func (r *FoodRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Food{}, "id = ?", id).Error
}

func (r *FoodRepo) List(dto foodDTO.FoodListDTO) (foods []models.Food, err error) {
	db := r.DB.Model(&models.Food{})

	if dto.Name != "" {
		db = db.Where("name ~ ?", dto.Name)
	}

	if dto.Description != "" {
		db = db.Where("description ~ ?", dto.Description)
	}

	if dto.Cost != 0 {
		db = db.Where("cost = ?", dto.Cost)
	}

	if dto.CategoryID != 0 {
		db = db.Where("category_id ~ ?", dto.CategoryID)
	}

	err = db.
		Order(dto.Sort).
		Preload("Category").
		Find(&foods).Error

	return foods, err
}
