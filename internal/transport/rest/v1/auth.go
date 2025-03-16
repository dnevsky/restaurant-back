package v1

import (
	userDto "github.com/dnevsky/restaurant-back/internal/dto/user"
	"github.com/dnevsky/restaurant-back/internal/pkg/config"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
	}
}

func (h *Handler) login(c *gin.Context) {
	var input userDto.AuthDto
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	resp, err := h.services.User.Login(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
		Data: tokenResponse{
			AccessToken: resp.AccessToken,
			ExpiresIn:   int64(config.Config.AccessTokenTTL.Minutes()),
		},
	})
}

func (h *Handler) register(c *gin.Context) {
	var input userDto.RegDto
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	resp, err := h.services.User.Register(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
		Data: tokenResponse{
			AccessToken: resp.AccessToken,
			ExpiresIn:   int64(config.Config.AccessTokenTTL.Minutes()),
		},
	})
}
