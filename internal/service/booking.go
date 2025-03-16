package service

import (
	"errors"
	bookingDTO "github.com/dnevsky/restaurant-back/internal/dto/booking"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/repository"
)

type Booking interface {
	Create(dto bookingDTO.BookingCreateDTO) (models.Booking, error)
	Update(dto bookingDTO.BookingUpdateDTO) error
	GetList(dto bookingDTO.BookingListDTO) ([]models.Booking, error)
	Get(id uint) (models.Booking, error)
	Delete(id uint) error
}

const bookingDefaultSort = "id ASC"

type BookingService struct {
	bookingRepo repository.BookingRepo
	tableRepo   repository.TableRepo
}

func NewBookingService(bookingRepo repository.BookingRepo, tableRepo repository.TableRepo) *BookingService {
	return &BookingService{
		bookingRepo: bookingRepo,
		tableRepo:   tableRepo,
	}
}

func (s *BookingService) Create(dto bookingDTO.BookingCreateDTO) (models.Booking, error) {
	table, err := s.tableRepo.FindWithStatus(dto.TableID, dto.Datetime)
	if err != nil {
		return models.Booking{}, errors.Join(models.ErrNotFound, errors.New("id_table"))
	}

	if !table.Available() {
		return models.Booking{}, errors.New("table is unavailable")
	}

	booking, err := s.bookingRepo.Create(models.NewBooking(dto.TableID, dto.Datetime, dto.Fullname, dto.Phone, dto.Email, dto.CountSeats, dto.NumberOfPeople))
	if err != nil {
		return models.Booking{}, err
	}

	return s.bookingRepo.Find(booking.ID)
}

func (s *BookingService) Update(dto bookingDTO.BookingUpdateDTO) error {
	booking, err := s.bookingRepo.Find(dto.ID)
	if err != nil {
		return err
	}

	if dto.TableID != nil {
		booking.TableID = *dto.TableID
	}

	if dto.Datetime != nil {
		booking.Datetime = *dto.Datetime
	}

	if dto.Fullname != nil {
		booking.FullName = *dto.Fullname
	}

	if dto.Phone != nil {
		booking.Phone = *dto.Phone
	}

	if dto.Email != nil {
		booking.Email = *dto.Email
	}

	if dto.CountSeats != nil {
		booking.CountSeats = *dto.CountSeats
	}

	if dto.NumberOfPeople != nil {
		booking.NumberOfPeople = *dto.NumberOfPeople
	}

	if dto.Status != nil {
		booking.Status = *dto.Status
	}

	return s.bookingRepo.Update(&booking)
}

func (s *BookingService) GetList(dto bookingDTO.BookingListDTO) ([]models.Booking, error) {
	if dto.Sort == "" {
		dto.Sort = bookingDefaultSort
	}

	return s.bookingRepo.List(dto)
}

func (s *BookingService) Get(id uint) (models.Booking, error) {
	return s.bookingRepo.Find(id)
}

func (s *BookingService) Delete(id uint) error {
	return s.bookingRepo.Delete(id)
}
