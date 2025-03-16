package postgres

import (
	tableDTO "github.com/dnevsky/restaurant-back/internal/dto/table"
	"github.com/dnevsky/restaurant-back/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TableRepo struct {
	DB *gorm.DB
}

func NewTableRepo(db *gorm.DB) *TableRepo {
	return &TableRepo{DB: db}
}

func (r *TableRepo) Create(table models.Table) (models.Table, error) {
	err := r.DB.Clauses(clause.Returning{}).Create(&table).Error
	return table, err
}

func (r *TableRepo) Update(table *models.Table) error {
	return r.DB.Save(table).Error
}

func (r *TableRepo) Find(id uint) (table models.Table, err error) {
	err = r.DB.First(&table, "id = ?", id).Error
	return table, err
}

func (r *TableRepo) FindWithStatus(id uint, dt time.Time) (table models.Table, err error) {
	if !dt.IsZero() {
		query := r.DB.Table("tables t").Select(
			"t.id, t.seat, t.number_seats, "+
				"CASE WHEN EXISTS ("+
				"  SELECT 1 FROM bookings b "+
				"  WHERE b.table_id = t.id "+
				"    AND b.status = ? "+
				"    AND ? >= b.datetime "+
				"    AND ? < (b.datetime + interval '1 hour')"+
				") THEN ? ELSE ? END as status",
			"approved", dt, dt, models.TableStatusBooked, models.TableStatusAvailable,
		)
		err = query.Where("t.id = ?", id).Scan(&table).Error
		return table, err
	}

	err = r.DB.First(&table, "id = ?", id).Error
	return table, err
}

func (r *TableRepo) FindBySeat(seat int) (table models.Table, err error) {
	err = r.DB.First(&table, "seat = ?", seat).Error
	return table, err
}

func (r *TableRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Table{}, "id = ?", id).Error
}

func (r *TableRepo) List(dto tableDTO.TableListDTO) (tables []models.Table, err error) {
	if !dto.Datetime.IsZero() {
		query := r.DB.Table("tables t").Select(
			"t.id, t.seat, t.number_seats, "+
				"CASE WHEN EXISTS ("+
				"  SELECT 1 FROM bookings b "+
				"  WHERE b.table_id = t.id "+
				"    AND b.status = ? "+
				"    AND ? >= b.datetime "+
				"    AND ? < (b.datetime + interval '1 hour')"+
				") THEN ? ELSE ? END as status",
			"approved", dto.Datetime, dto.Datetime,
			models.TableStatusBooked, models.TableStatusAvailable,
		)

		if dto.Seat != 0 {
			query = query.Where("t.seat = ?", dto.Seat)
		}

		if dto.NumberSeats != 0 {
			query = query.Where("t.number_seats = ?", dto.NumberSeats)
		}

		if dto.Status != "" {
			query = query.Having("status = ?", dto.Status)
		}
		query = query.Order(dto.Sort)
		err = query.Scan(&tables).Error
		return tables, err
	}

	db := r.DB.Model(&models.Table{})

	if dto.Seat != 0 {
		db = db.Where("seat = ?", dto.Seat)
	}

	if dto.NumberSeats != 0 {
		db = db.Where("number_seats = ?", dto.NumberSeats)
	}

	if dto.Status != "" {
		db = db.Where("status = ?", dto.Status)
	}

	err = db.
		Order(dto.Sort).
		Find(&tables).Error

	return tables, err
}
