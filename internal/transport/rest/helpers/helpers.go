package helpers

import (
	"github.com/dnevsky/restaurant-back/internal/repository"
	"github.com/gin-gonic/gin"
)

type Helpers interface {
	GetUserIdAuthorization(c *gin.Context) (uint, error)
	GetAdminIdAuthorization(c *gin.Context) (uint, error)
	BindData(c *gin.Context, req interface{}) error
	ErrorsHandle(c *gin.Context, err error)
	LogError(err error)
	GetIdFromPath(c *gin.Context, key string) (uint, error)
}

type Manager struct {
	UserRepository repository.UserRepo
}

func NewManager(
	userRepo repository.UserRepo,
) *Manager {
	return &Manager{
		UserRepository: userRepo,
	}
}
