package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type FoodUpdateDTO struct {
	dto.ServiceDTO
	ID          uint
	Name        *string `json:"name,omitempty" form:"name" validate:"omitempty,min=3,max=255"`
	Description *string `json:"description,omitempty" form:"description" validate:"omitempty,min=3,max=255"`
	Cost        *int    `json:"cost,omitempty" form:"cost" validate:"omitempty,min=1"`
	CategoryID  *uint   `json:"category_id,omitempty" form:"category_id" validate:"omitempty"`
	Picture     *string `json:"picture,omitempty" form:"picture" validate:"omitempty,url"`
}

func (dto *FoodUpdateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
