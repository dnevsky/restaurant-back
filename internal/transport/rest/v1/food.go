package v1

import (
	foodDTO "github.com/dnevsky/restaurant-back/internal/dto/food"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/middleware"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initFoodRoutes(api *gin.RouterGroup) {
	auth := api.Group("/", middleware.AuthUser(h.services.TokenManager))
	{
		foodGroup := auth.Group("/food")
		{
			foodGroup.POST("/", h.foodCreate)
			foodGroup.GET("/:id", h.foodGet)
			foodGroup.DELETE("/:id", h.foodDelete)
			foodGroup.PUT("/:id", h.foodUpdate)
		}
	}
	foodNoAuthGroup := api.Group("/food")
	{
		foodNoAuthGroup.GET("/", h.foodGetList)
	}
}

func (h *Handler) foodCreate(c *gin.Context) {
	var input foodDTO.FoodCreateDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if !input.AuthUser.IsAdmin() {
		h.helpers.ErrorsHandle(c, models.ErrAccessDenied)
		return
	}

	items, err := h.services.Food.Create(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
		Code: http.StatusOK,
	})
}

func (h *Handler) foodGet(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	item, err := h.services.Food.Get(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: item,
	})
}

func (h *Handler) foodDelete(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	_, err = h.helpers.GetAdminIdAuthorization(c)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	err = h.services.Food.Delete(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) foodUpdate(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	var input foodDTO.FoodUpdateDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}
	input.ID = id

	if err = input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if !input.AuthUser.IsAdmin() {
		h.helpers.ErrorsHandle(c, models.ErrAccessDenied)
		return
	}

	err = h.services.Food.Update(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) foodGetList(c *gin.Context) {
	var input foodDTO.FoodListDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	items, err := h.services.Food.GetList(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
	})
}
