package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type CategoryUpdateDTO struct {
	dto.ServiceDTO
	ID   uint
	Name *string `json:"name,omitempty" form:"name" validate:"omitempty,min=3,max=255"`
}

func (dto *CategoryUpdateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
