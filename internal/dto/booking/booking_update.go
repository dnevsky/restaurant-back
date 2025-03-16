package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
	"time"
)

type BookingUpdateDTO struct {
	dto.ServiceDTO
	ID             uint
	TableID        *uint                 `json:"table_id,omitempty" form:"table_id" validate:"omitempty"`
	Datetime       *time.Time            `json:"datetime,omitempty" form:"datetime" validate:"omitempty"`
	Fullname       *string               `json:"fullname,omitempty" form:"fullname" validate:"omitempty,min=3,max=255"`
	Phone          *string               `json:"phone,omitempty" form:"phone" validate:"omitempty,min=11,max=13"`
	Email          *string               `json:"email,omitempty" form:"email" validate:"omitempty,email"`
	CountSeats     *int                  `json:"count_seats,omitempty" form:"count_seats" validate:"omitempty,min=1"`
	NumberOfPeople *int                  `json:"number_of_people,omitempty" form:"number_of_people" validate:"omitempty,min=1"`
	Status         *models.BookingStatus `json:"status,omitempty" form:"status" validate:"omitempty,oneof=new approved rejected finished"`
}

func (dto *BookingUpdateDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
