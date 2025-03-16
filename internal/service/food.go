package service

import (
	"errors"
	foodDTO "github.com/dnevsky/restaurant-back/internal/dto/food"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/repository"
)

type Food interface {
	Create(dto foodDTO.FoodCreateDTO) (models.Food, error)
	Update(dto foodDTO.FoodUpdateDTO) error
	GetList(dto foodDTO.FoodListDTO) ([]models.Food, error)
	Get(id uint) (models.Food, error)
	Delete(id uint) error
}

const foodDefaultSort = "category_id,name ASC"

type FoodService struct {
	foodRepo     repository.FoodRepo
	categoryRepo repository.CategoryRepo
}

func NewFoodService(foodRepo repository.FoodRepo, categoryRepo repository.CategoryRepo) *FoodService {
	return &FoodService{
		foodRepo:     foodRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *FoodService) Create(dto foodDTO.FoodCreateDTO) (models.Food, error) {
	if _, err := s.categoryRepo.Find(dto.CategoryID); err != nil {
		return models.Food{}, errors.Join(models.ErrNotFound, errors.New("category_id"))
	}
	return s.foodRepo.Create(models.NewFood(dto.Name, dto.Description, dto.Cost, dto.CategoryID, dto.Picture))
}

func (s *FoodService) Update(dto foodDTO.FoodUpdateDTO) error {
	food, err := s.foodRepo.Find(dto.ID)
	if err != nil {
		return err
	}

	if dto.Name != nil {
		food.Name = *dto.Name
	}

	if dto.Description != nil {
		food.Description = *dto.Description
	}

	if dto.Cost != nil {
		food.Cost = *dto.Cost
	}

	if dto.CategoryID != nil {
		food.CategoryID = *dto.CategoryID
	}

	if dto.Picture != nil {
		food.Picture = *dto.Picture
	}

	return s.foodRepo.Update(&food)
}

func (s *FoodService) GetList(dto foodDTO.FoodListDTO) ([]models.Food, error) {
	if dto.Sort == "" {
		dto.Sort = foodDefaultSort
	}

	return s.foodRepo.List(dto)
}

func (s *FoodService) Get(id uint) (models.Food, error) {
	return s.foodRepo.Find(id)
}

func (s *FoodService) Delete(id uint) error {
	return s.foodRepo.Delete(id)
}
