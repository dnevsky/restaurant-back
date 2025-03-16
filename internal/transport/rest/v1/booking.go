package v1

import (
	bookingDTO "github.com/dnevsky/restaurant-back/internal/dto/booking"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/middleware"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initBookingRoutes(api *gin.RouterGroup) {
	auth := api.Group("/", middleware.AuthUser(h.services.TokenManager))
	{
		bookingGroup := auth.Group("/booking")
		{
			bookingGroup.GET("/:id", h.bookingGet)
			bookingGroup.DELETE("/:id", h.bookingDelete)
			bookingGroup.PUT("/:id", h.bookingUpdate)
			bookingGroup.GET("/", h.bookingGetList)

		}
	}
	bookingNoAuthGroup := api.Group("/booking")
	{
		bookingNoAuthGroup.POST("/", h.bookingCreate)

	}
}

func (h *Handler) bookingCreate(c *gin.Context) {
	var input bookingDTO.BookingCreateDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	items, err := h.services.Booking.Create(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
		Code: http.StatusOK,
	})
}

func (h *Handler) bookingGet(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	item, err := h.services.Booking.Get(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: item,
	})
}

func (h *Handler) bookingDelete(c *gin.Context) {
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

	err = h.services.Booking.Delete(id)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) bookingUpdate(c *gin.Context) {
	id, err := h.helpers.GetIdFromPath(c, "id")
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	var input bookingDTO.BookingUpdateDTO
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

	err = h.services.Booking.Update(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Code: http.StatusOK,
	})
}

func (h *Handler) bookingGetList(c *gin.Context) {
	var input bookingDTO.BookingListDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	items, err := h.services.Booking.GetList(input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: items,
	})
}
