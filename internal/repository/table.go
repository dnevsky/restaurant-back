package repository

import (
	tableDTO "github.com/dnevsky/restaurant-back/internal/dto/table"
	"github.com/dnevsky/restaurant-back/internal/models"
	"time"
)

type TableRepo interface {
	Create(table models.Table) (models.Table, error)
	Update(table *models.Table) error
	Find(id uint) (table models.Table, err error)
	FindWithStatus(id uint, dt time.Time) (table models.Table, err error)
	FindBySeat(seat int) (table models.Table, err error)
	Delete(id uint) error
	List(dto tableDTO.TableListDTO) (tables []models.Table, err error)
}
