package middleware

import (
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/pkg/auth"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/helpers"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthUser(tokenManager auth.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseAuthHeader(c, tokenManager)
		if err != nil {
			response.NewResponse(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		c.Set(helpers.UserCtx, id)
	}
}

func parseAuthHeader(c *gin.Context, tokenManager auth.TokenManager) (string, error) {
	header := c.GetHeader(helpers.AuthorizationHeader)
	if header == "" {
		return "", models.ErrEmptyAuthHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", models.ErrInvalidAuthHeader
	}
	if len(headerParts[1]) == 0 {
		return "", models.ErrTokenIsEmpty
	}

	return tokenManager.ParseAccessToken(headerParts[1])
}
