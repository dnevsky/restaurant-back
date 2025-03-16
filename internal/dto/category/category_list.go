package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type CategoryListDTO struct {
	dto.ServiceDTO
	Sort string `form:"sort" json:"sort"`
	Name string `json:"name" form:"name" validate:"omitempty,min=3,max=255"`
}

func (dto *CategoryListDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
