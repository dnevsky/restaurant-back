package v1

import (
	categoryDTO "github.com/dnevsky/restaurant-back/internal/dto/category"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/middleware"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initCategoryRoutes(api *gin.RouterGroup) {
	auth := api.Group("/", middleware.AuthUser(h.services.TokenManager))
	{
		categoryGroup := auth.Group("/category")
		{
			categoryGroup.POST("/", h.categoryCreate)
			categoryGroup.GET("/", h.categoryGetList)
			categoryGroup.GET("/:id", h.categoryGet)
			categoryGroup.DELETE("/:id", h.categoryDelete)
			categoryGroup.PUT("/:id", h.categoryUpdate)
		}
	}
}

func (h *Handler) categoryCreate(c *gin.Context) {
	var input categoryDTO.CategoryCreateDTO
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

	items, err := h.services.Category.Create(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
		Code: http.StatusOK,
	})
}

func (h *Handler) categoryGet(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	item, err := h.services.Category.Get(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: item,
	})
}

func (h *Handler) categoryDelete(c *gin.Context) {
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

	err = h.services.Category.Delete(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) categoryUpdate(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	var input categoryDTO.CategoryUpdateDTO
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

	err = h.services.Category.Update(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) categoryGetList(c *gin.Context) {
	var input categoryDTO.CategoryListDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	items, err := h.services.Category.GetList(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
	})
}
