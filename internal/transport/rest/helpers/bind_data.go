package helpers

import (
	"github.com/dnevsky/restaurant-back/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
)

const (
	authUserField = "AuthUser"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func (m *Manager) BindData(c *gin.Context, req interface{}) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}

	if err := c.Bind(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			return models.ErrInvalidRequestParams
		}

		return err
	}

	err = m.BindAuthUser(c, &req)
	if err != nil {
		return models.ErrUnauthorized
	}

	return nil
}

func (m *Manager) BindAuthUser(c *gin.Context, req *interface{}) error {
	uid, err := m.GetUserIdAuthorization(c)
	if err == nil {
		user, err := m.UserRepository.Find(uid)
		if err != nil {
			return models.ErrUnauthorized
		}
		v := reflect.ValueOf(*req).Elem()

		if f := v.FieldByName(authUserField); f.IsValid() {
			f.Set(reflect.ValueOf(&user))
		}
	}
	return nil
}

func (m *Manager) GetIdFromPath(c *gin.Context, key string) (uint, error) {
	param := c.Param(key)
	intParam, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return uint(intParam), nil
}
