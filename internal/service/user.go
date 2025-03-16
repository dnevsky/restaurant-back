package service

import (
	userDto "github.com/dnevsky/restaurant-back/internal/dto/user"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/pkg/auth"
	"github.com/dnevsky/restaurant-back/internal/pkg/config"
	"github.com/dnevsky/restaurant-back/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type User interface {
	Login(req userDto.AuthDto) (session models.Session, err error)
	Register(req userDto.RegDto) (session models.Session, err error)
}

type UserService struct {
	userRepo     repository.UserRepo
	tokenManager auth.TokenManager
}

func NewUserService(userRepo repository.UserRepo, tokenManager auth.TokenManager) *UserService {
	return &UserService{
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (s *UserService) Login(req userDto.AuthDto) (session models.Session, err error) {
	var user *models.User

	user, err = s.authEmail(req)
	if err != nil {
		return models.Session{}, err
	}

	if user == nil || user.ID == 0 {
		return models.Session{}, nil
	}

	if !user.IsAdmin() {
		return models.Session{}, models.ErrAccessDenied
	}

	tokens, err := s.createTokens(user)
	if err != nil {
		return models.Session{}, err
	}

	return tokens, nil
}

func (s *UserService) Register(req userDto.RegDto) (session models.Session, err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.NewUser(req.Name, req.Email, string(hashedPassword))

	user, err = s.userRepo.Create(user)
	if err != nil {
		return models.Session{}, models.ErrAlreadyExists
	}

	tokens, err := s.createTokens(user)
	if err != nil {
		return models.Session{}, err
	}

	return tokens, nil
}

func (s *UserService) authEmail(req userDto.AuthDto) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(req.Email)

	if err != nil {
		return nil, models.ErrInvalidAuthCreds
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, models.ErrInvalidAuthCreds
	}

	return user, err
}

func (s *UserService) createTokens(user *models.User) (token models.Session, err error) {
	accessToken, err := s.tokenManager.NewAccessToken(strconv.Itoa(int(user.ID)), config.Config.AccessTokenTTL)
	if err != nil {
		return models.Session{}, err
	}

	return models.Session{
		AccessToken: accessToken,
	}, nil
}
