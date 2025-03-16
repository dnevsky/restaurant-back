package postgres

import (
	"github.com/dnevsky/restaurant-back/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Find(id uint) (user models.User, err error) {
	err = r.DB.
		Where("id = ?", id).
		First(&user).Error
	return user, err
}

func (r *UserRepo) Create(user *models.User) (*models.User, error) {
	err := r.DB.Clauses(clause.Returning{}).Create(&user).Error
	return user, err
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "lower(email) = ?", email).Error
	return &user, err
}

func (r *UserRepo) Update(user *models.User) error {
	return r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Select("*").Updates(&user).Error
}
