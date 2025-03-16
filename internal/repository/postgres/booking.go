package postgres

import (
	bookingDTO "github.com/dnevsky/restaurant-back/internal/dto/booking"
	"github.com/dnevsky/restaurant-back/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRepo struct {
	DB *gorm.DB
}

func NewBookingRepo(db *gorm.DB) *BookingRepo {
	return &BookingRepo{DB: db}
}

func (r *BookingRepo) Create(booking models.Booking) (models.Booking, error) {
	err := r.DB.Clauses(clause.Returning{}).Create(&booking).Error
	return booking, err
}

func (r *BookingRepo) Update(booking *models.Booking) error {
	return r.DB.Omit("Table").Save(booking).Error
}

func (r *BookingRepo) Find(id uint) (booking models.Booking, err error) {
	err = r.DB.Preload("Table", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).First(&booking, "id = ?", id).Error
	return booking, err
}

func (r *BookingRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Booking{}, "id = ?", id).Error
}

func (r *BookingRepo) List(dto bookingDTO.BookingListDTO) (bookings []models.Booking, err error) {
	db := r.DB.Model(&models.Booking{})

	if dto.TableID != 0 {
		db = db.Where("table_id = ?", dto.TableID)
	}

	if !dto.Datetime.IsZero() {
		db = db.Where("datetime = ?", dto.Datetime)
	}

	if dto.Fullname != "" {
		db = db.Where("fullname = ?", dto.Fullname)
	}

	if dto.Phone != "" {
		db = db.Where("phone = ?", dto.Phone)
	}

	if dto.Email != "" {
		db = db.Where("email = ?", dto.Email)
	}

	if dto.CountSeats != 0 {
		db = db.Where("count_seats = ?", dto.CountSeats)
	}

	if dto.NumberOfPeople != 0 {
		db = db.Where("number_of_people = ?", dto.NumberOfPeople)
	}

	err = db.
		Order(dto.Sort).
		Preload("Table", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Find(&bookings).Error

	return bookings, err
}
