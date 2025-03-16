package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type TableCreateDTO struct {
	dto.ServiceDTO
	Seat        int `json:"seat" form:"seat" validate:"required,min=1"`
	NumberSeats int `json:"number_seats" form:"number_seats" validate:"required,min=1"`
}

func (dto *TableCreateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
