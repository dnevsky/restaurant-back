package v1

import (
	"github.com/dnevsky/restaurant-back/internal/service"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/helpers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	helpers  helpers.Helpers
}

func NewHandler(services *service.Service, helpers helpers.Helpers) *Handler {
	return &Handler{
		services: services,
		helpers:  helpers,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initBookingRoutes(v1)
		h.initCategoryRoutes(v1)
		h.initFoodRoutes(v1)
		h.initTableRoutes(v1)
		h.initAuthRoutes(v1)
	}
}
