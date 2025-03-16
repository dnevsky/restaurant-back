package service

import (
	categoryDTO "github.com/dnevsky/restaurant-back/internal/dto/category"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/repository"
)

type Category interface {
	Create(dto categoryDTO.CategoryCreateDTO) (models.Category, error)
	Update(dto categoryDTO.CategoryUpdateDTO) error
	GetList(dto categoryDTO.CategoryListDTO) ([]models.Category, error)
	Get(id uint) (models.Category, error)
	Delete(id uint) error
}

const categoryDefaultSort = "id ASC"

type CategoryService struct {
	categoryRepo repository.CategoryRepo
}

func NewCategoryService(categoryRepo repository.CategoryRepo) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) Create(dto categoryDTO.CategoryCreateDTO) (models.Category, error) {
	return s.categoryRepo.Create(models.NewCategory(dto.Name))
}

func (s *CategoryService) Update(dto categoryDTO.CategoryUpdateDTO) error {
	category, err := s.categoryRepo.Find(dto.ID)
	if err != nil {
		return err
	}

	if dto.Name != nil {
		category.Name = *dto.Name
	}

	return s.categoryRepo.Update(&category)
}

func (s *CategoryService) GetList(dto categoryDTO.CategoryListDTO) ([]models.Category, error) {
	if dto.Sort == "" {
		dto.Sort = categoryDefaultSort
	}

	return s.categoryRepo.List(dto)
}

func (s *CategoryService) Get(id uint) (models.Category, error) {
	return s.categoryRepo.Find(id)
}

func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}
