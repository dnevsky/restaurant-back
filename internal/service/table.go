package service

import (
	tableDTO "github.com/dnevsky/restaurant-back/internal/dto/table"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/repository"
)

type Table interface {
	Create(dto tableDTO.TableCreateDTO) (models.Table, error)
	Update(dto tableDTO.TableUpdateDTO) error
	GetList(dto tableDTO.TableListDTO) ([]models.Table, error)
	Get(id uint) (models.Table, error)
	Delete(id uint) error
}

const tableDefaultSort = "seat ASC"

type TableService struct {
	tableRepo repository.TableRepo
}

func NewTableService(tableRepo repository.TableRepo) *TableService {
	return &TableService{tableRepo: tableRepo}
}

func (s *TableService) Create(dto tableDTO.TableCreateDTO) (models.Table, error) {
	if _, err := s.tableRepo.FindBySeat(dto.Seat); err == nil {
		return models.Table{}, models.ErrAlreadyExists
	}
	return s.tableRepo.Create(models.NewTable(dto.Seat, dto.NumberSeats))
}

func (s *TableService) Update(dto tableDTO.TableUpdateDTO) error {
	table, err := s.tableRepo.Find(dto.ID)
	if err != nil {
		return err
	}

	if dto.Seat != nil {
		if _, err = s.tableRepo.FindBySeat(*dto.Seat); err == nil {
			return models.ErrAlreadyExists
		}
		table.Seat = *dto.Seat
	}

	if dto.NumberSeats != nil {
		table.NumberSeats = *dto.NumberSeats
	}

	return s.tableRepo.Update(&table)
}

func (s *TableService) GetList(dto tableDTO.TableListDTO) ([]models.Table, error) {
	if dto.Sort == "" {
		dto.Sort = tableDefaultSort
	}

	return s.tableRepo.List(dto)
}

func (s *TableService) Get(id uint) (models.Table, error) {
	return s.tableRepo.Find(id)
}

func (s *TableService) Delete(id uint) error {
	return s.tableRepo.Delete(id)
}
