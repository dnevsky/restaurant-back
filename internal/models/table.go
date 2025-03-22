package models

type Table struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Seat        int          `json:"seat"`
	NumberSeats int          `json:"number_seats"`
	Status      *TableStatus `json:"status" gorm:"-"`
	Bookings    []Booking    `json:"bookings,omitempty" gorm:"foreignKey:TableID"`
}

func (t Table) Available() bool {
	if t.Status == nil {
		return false
	}
	return *t.Status == TableStatusAvailable
}

func NewTable(seat int, numberSeats int) Table {
	return Table{
		Seat:        seat,
		NumberSeats: numberSeats,
	}
}

type TableStatus string

const (
	TableStatusAvailable   TableStatus = "available"
	TableStatusUnavailable TableStatus = "unavailable"
	TableStatusBooked      TableStatus = "booked"
)
