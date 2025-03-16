package helpers

import (
	"errors"
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/dnevsky/restaurant-back/internal/pkg/logger"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (m *Manager) getErrCodeType(err error) (int, models.ErrType) {
	// TODO: варианты ошибок будут пополняться

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, models.ErrTypeDefault
	}

	if errors.Is(err, models.ErrAccessDenied) {
		return http.StatusForbidden, models.ErrTypeDefault
	}

	if strings.Contains(err.Error(), models.ErrBrokenPipe.Error()) ||
		strings.Contains(err.Error(), models.ErrConnectionResetByPeer.Error()) {
		return http.StatusBadRequest, models.ErrTypeDefault
	}

	if errors.Is(err, models.ErrUnauthorized) {
		return http.StatusUnauthorized, models.ErrTypeDefault
	}

	if errors.Is(err, models.ErrBadRequest) {
		return http.StatusBadRequest, models.ErrTypeDefault
	}

	if errors.Is(err, models.ErrAlreadyExists) {
		return http.StatusConflict, models.ErrTypeDefault
	}

	// bad request ошибки
	// if errors.Is(err, aeroflot.ErrClassUpgradeDisallow) {
	// 	return http.StatusBadRequest, models.ErrTypeDefault
	// }

	var validationErr validator.ValidationErrors

	if errors.As(err, &validationErr) {
		return http.StatusBadRequest, models.ErrTypeValidation
	} else {
		return http.StatusInternalServerError, models.ErrTypeDefault
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Minimum string length is " + fe.Param()
	case "max":
		return "Maximum string length is " + fe.Param()
	case "oneof":
		return "Field can be one of: " + fe.Param()
	case "noSpecialChars":
		return "Bad characters into tag"
	case "url":
		return "Bad url format"
	case "phone":
		return "Bad phone format"
	case "email":
		return "Bad email format"
	}
	return "Unknown error"
}

func (m *Manager) canLogError(code int) bool {
	if code >= 500 {
		return true
	}
	return false
}

func (m *Manager) defaultValidationErrorsHandle(c *gin.Context, err error, errCode int) {
	response.JsonResponse(c.Writer, response.Data{
		Code: errCode,
		Text: err.Error(),
	})
}

func (m *Manager) validationErrorsHandle(c *gin.Context, err error, errCode int) {
	var validationErr validator.ValidationErrors

	if errors.As(err, &validationErr) {
		out := make([]ErrorMsg, len(validationErr))
		for i, vErr := range err.(validator.ValidationErrors) {
			out[i] = ErrorMsg{vErr.Field(), getErrorMsg(vErr)}
		}
		response.JsonResponse(c.Writer, response.Data{
			Code:         http.StatusBadRequest,
			ClientErrors: out,
		})
	} else {
		m.defaultValidationErrorsHandle(c, err, errCode)
	}
}

func (m *Manager) customValidationErrorsHandle(c *gin.Context, err error, errCode int) {
	// if errors.Is(err, models.ErrNoColumnsToBuildFile) {
	// 	response.JsonResponse(c.Writer, response.Data{
	// 		Code: http.StatusBadRequest,
	// 		ClientErrors: ErrorMsg{
	// 			Field:   "columns",
	// 			Message: err.Error(),
	// 		},
	// 	})
	// } else {
	m.defaultValidationErrorsHandle(c, err, errCode)
	// }
}

func (m *Manager) LogError(err error) {
	errCode, _ := m.getErrCodeType(err)
	if m.canLogError(errCode) {
		logger.Log.Error(err)
	}
}

func (m *Manager) ErrorsHandle(c *gin.Context, err error) {
	m.LogError(err)

	errCode, errType := m.getErrCodeType(err)
	switch errType {
	case models.ErrTypeDefault:
		m.defaultValidationErrorsHandle(c, err, errCode)
	case models.ErrTypeValidation:
		m.validationErrorsHandle(c, err, errCode)
	case models.ErrTypeCustomValidation:
		m.customValidationErrorsHandle(c, err, errCode)
	default:
	}
}
