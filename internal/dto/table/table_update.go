package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type TableUpdateDTO struct {
	dto.ServiceDTO
	ID          uint
	Seat        *int                `json:"seat,omitempty" form:"seat" validate:"omitempty,min=1"`
	NumberSeats *int                `json:"number_seats,omitempty" form:"number_seats" validate:"omitempty,min=1"`
	Status      *models.TableStatus `json:"status,omitempty" form:"status" validate:"omitempty,oneof=available unavailable booked"`
}

func (dto *TableUpdateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
