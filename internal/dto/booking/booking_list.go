package table

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
	"time"
)

type BookingListDTO struct {
	dto.ServiceDTO
	Sort           string               `form:"sort" json:"sort"`
	TableID        uint                 `json:"table_id" form:"table_id" validate:"omitempty"`
	Datetime       time.Time            `json:"datetime" form:"datetime" validate:"omitempty"`
	Fullname       string               `json:"fullname" form:"fullname" validate:"omitempty,min=3,max=255"`
	Phone          string               `json:"phone" form:"phone" validate:"omitempty,min=11,max=13"`
	Email          string               `json:"email" form:"email" validate:"omitempty,email"`
	CountSeats     int                  `json:"count_seats" form:"count_seats" validate:"omitempty,min=1"`
	NumberOfPeople int                  `json:"number_of_people" form:"number_of_people" validate:"omitempty,min=1"`
	Status         models.BookingStatus `json:"status" form:"status" validate:"omitempty,oneof=new approved rejected finished"`
}

func (dto *BookingListDTO) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
