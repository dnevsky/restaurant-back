package helpers

import (
	"github.com/dnevsky/restaurant-back/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"

	UserCtx = "userId"
)

func (m *Manager) GetUserIdAuthorization(c *gin.Context) (uint, error) {
	if userId, exists := c.Get(UserCtx); exists {
		userId := userId.(string)

		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			return 0, err
		}
		return uint(userIdInt), nil
	}
	return 0, models.ErrUnauthorized
}

func (m *Manager) GetAdminIdAuthorization(c *gin.Context) (uint, error) {
	userId, err := m.GetUserIdAuthorization(c)
	if err != nil {
		return 0, err
	}

	user, err := m.UserRepository.Find(userId)
	if err != nil {
		return 0, err
	}

	if user.IsAdmin() {
		return userId, nil
	}

	return 0, models.ErrAccessDenied
}
