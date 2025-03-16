package repository

import (
	bookingDTO "github.com/dnevsky/restaurant-back/internal/dto/booking"
	"github.com/dnevsky/restaurant-back/internal/models"
)

type BookingRepo interface {
	Create(booking models.Booking) (models.Booking, error)
	Update(booking *models.Booking) error
	Find(id uint) (booking models.Booking, err error)
	Delete(id uint) error
	List(dto bookingDTO.BookingListDTO) (bookings []models.Booking, err error)
}
