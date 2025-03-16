package user

import (
	"github.com/dnevsky/restaurant-back/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type RegDto struct {
	dto.ServiceDTO
	Email    string `json:"email" form:"email" conform:"email,lower" validate:"required,min=3,max=255,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
	Name     string `json:"name" form:"name" validate:"required,min=3,max=255"`
}

func (dto *RegDto) Validate() error {
	validate := validator.New()
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
