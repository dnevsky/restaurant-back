package models

import "time"

type Booking struct {
	ID             uint          `json:"id" gorm:"primaryKey"`
	TableID        uint          `json:"table_id"`
	Table          Table         `json:"table" gorm:"foreignKey:TableID;"`
	Datetime       time.Time     `json:"datetime"`
	FullName       string        `json:"fullname" gorm:"column:fullname"`
	Phone          string        `json:"phone"`
	Email          string        `json:"email"`
	CountSeats     int           `json:"count_seats"`
	NumberOfPeople int           `json:"number_of_people"`
	Status         BookingStatus `json:"status"`
}

func NewBooking(tableId uint, datetime time.Time, fullname, phone, email string, countSeats, numberOfPeople int) Booking {
	return Booking{
		TableID:        tableId,
		Datetime:       datetime,
		FullName:       fullname,
		Phone:          phone,
		Email:          email,
		CountSeats:     countSeats,
		NumberOfPeople: numberOfPeople,
		Status:         BookingStatusApproved,
	}
}

type BookingStatus string

const (
	BookingStatusNew      BookingStatus = "new"      // новая бронь
	BookingStatusApproved BookingStatus = "approved" // бронь подтверждена
	BookingStatusRejected BookingStatus = "rejected" // бронь не подтверждена
	BookingStatusFinished BookingStatus = "finished" // бронью воспользовались
)
