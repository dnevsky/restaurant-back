package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type FoodCreateDTO struct {
	dto.ServiceDTO
	Name        string `json:"name" form:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" form:"description" validate:"required,min=3,max=255"`
	Cost        int    `json:"cost" form:"cost" validate:"required,min=1"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
	Picture     string `json:"picture" form:"picture" validate:"required,url"`
}

func (dto *FoodCreateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
