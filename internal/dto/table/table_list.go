package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
	"time"
)

type TableListDTO struct {
	dto.ServiceDTO
	Sort        string             `form:"sort" json:"sort"`
	Datetime    time.Time          `form:"datetime" json:"datetime" validate:"required"`
	Seat        int                `json:"seat" form:"seat" validate:"omitempty,min=1"`
	NumberSeats int                `json:"number_seats" form:"number_seats" validate:"omitempty,min=1"`
	Status      models.TableStatus `json:"status" form:"status" validate:"omitempty,oneof=available unavailable booked"`
}

func (dto *TableListDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
