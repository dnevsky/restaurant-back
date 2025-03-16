package repository

import (
	"github.com/dnevsky/restaurant-back/internal/repository/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB

	UserRepo     UserRepo
	BookingRepo  BookingRepo
	CategoryRepo CategoryRepo
	FoodRepo     FoodRepo
	TableRepo    TableRepo
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB:           db,
		UserRepo:     postgres.NewUserRepo(db),
		BookingRepo:  postgres.NewBookingRepo(db),
		CategoryRepo: postgres.NewCategoryRepo(db),
		FoodRepo:     postgres.NewFoodRepo(db),
		TableRepo:    postgres.NewTableRepo(db),
	}
}
