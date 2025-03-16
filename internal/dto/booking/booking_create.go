package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
	"time"
)

type BookingCreateDTO struct {
	dto.ServiceDTO
	TableID        uint      `json:"table_id" form:"table_id" validate:"required"`
	Datetime       time.Time `json:"datetime" form:"datetime" validate:"required"`
	Fullname       string    `json:"fullname" form:"fullname" validate:"required,min=3,max=255"`
	Phone          string    `json:"phone" form:"phone" validate:"required,min=11,max=13"`
	Email          string    `json:"email" form:"email" validate:"required,email"`
	CountSeats     int       `json:"count_seats" form:"count_seats" validate:"required,min=1"`
	NumberOfPeople int       `json:"number_of_people" form:"number_of_people" validate:"required,min=1"`
}

func (dto *BookingCreateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
