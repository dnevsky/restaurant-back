package v1

import (
	tableDTO "github.com/dnevsky/restaurant-back/internal/dto/table"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/middleware"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initTableRoutes(api *gin.RouterGroup) {
	auth := api.Group("/", middleware.AuthUser(h.services.TokenManager))
	{
		tableGroup := auth.Group("/table")
		{
			tableGroup.POST("/", h.tableCreate)
			tableGroup.GET("/:id", h.tableGet)
			tableGroup.DELETE("/:id", h.tableDelete)
			tableGroup.PUT("/:id", h.tableUpdate)
		}
	}
	tableNoAuthGroup := api.Group("/table")
	{
		tableNoAuthGroup.GET("/", h.tableGetList)
	}
}

func (h *Handler) tableCreate(c *gin.Context) {
	var input tableDTO.TableCreateDTO
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

	items, err := h.services.Table.Create(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
		Code: http.StatusOK,
	})
}

func (h *Handler) tableGet(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	item, err := h.services.Table.Get(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: item,
	})
}

func (h *Handler) tableDelete(c *gin.Context) {
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

	err = h.services.Table.Delete(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) tableUpdate(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	var input tableDTO.TableUpdateDTO
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

	err = h.services.Table.Update(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) tableGetList(c *gin.Context) {
	var input tableDTO.TableListDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	items, err := h.services.Table.GetList(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
	})
}
