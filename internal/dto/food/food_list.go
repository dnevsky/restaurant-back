package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type FoodListDTO struct {
	dto.ServiceDTO
	Sort        string `form:"sort" json:"sort"`
	Name        string `json:"name" form:"name" validate:"omitempty,min=3,max=255"`
	Description string `json:"description" form:"description" validate:"omitempty,min=3,max=255"`
	Cost        int    `json:"cost" form:"cost" validate:"omitempty,min=1"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"omitempty"`
}

func (dto *FoodListDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
