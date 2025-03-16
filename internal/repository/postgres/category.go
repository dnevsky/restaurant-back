package postgres

import (
	categoryDTO "github.com/dnevsky/restaurant-back/internal/dto/category"
	"github.com/dnevsky/restaurant-back/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) Create(category models.Category) (models.Category, error) {
	err := r.DB.Clauses(clause.Returning{}).Create(&category).Error
	return category, err
}

func (r *CategoryRepo) Update(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepo) Find(id uint) (category models.Category, err error) {
	err = r.DB.First(&category, "id = ?", id).Error
	return category, err
}

func (r *CategoryRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Category{}, "id = ?", id).Error
}

func (r *CategoryRepo) List(dto categoryDTO.CategoryListDTO) (categories []models.Category, err error) {
	db := r.DB.Model(&models.Category{})

	if dto.Name != "" {
		db = db.Where("name ~ ?", dto.Name)
	}

	err = db.
		Order(dto.Sort).
		Find(&categories).Error

	return categories, err
}
