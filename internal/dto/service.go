package dto

import "github.com/dnevsky/restaurant-back/internal/models"

type ServiceDTO struct {
	AuthUser *models.User
}

type PaginationDTO struct {
	Limit  int `form:"limit,default=50" json:"limit"`
	Page   int `form:"page,default=1" json:"page"`
	Offset int
}
