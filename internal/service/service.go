package service

import (
	"github.com/dnevsky/restaurant-back/internal/pkg/auth"
	"github.com/dnevsky/restaurant-back/internal/repository"
)

type Service struct {
	TokenManager auth.TokenManager
	Booking      Booking
	Category     Category
	Food         Food
	Table        Table
	User         User
}

type Deps struct {
	Repository   *repository.Repository
	TokenManager auth.TokenManager
}

func NewService(deps Deps) (*Service, error) {
	bookingService := NewBookingService(deps.Repository.BookingRepo, deps.Repository.TableRepo)
	categoryService := NewCategoryService(deps.Repository.CategoryRepo)
	foodService := NewFoodService(deps.Repository.FoodRepo, deps.Repository.CategoryRepo)
	tableService := NewTableService(deps.Repository.TableRepo)
	userService := NewUserService(deps.Repository.UserRepo, deps.TokenManager)

	return &Service{
		TokenManager: deps.TokenManager,
		Booking:      bookingService,
		Category:     categoryService,
		Food:         foodService,
		Table:        tableService,
		User:         userService,
	}, nil
}
