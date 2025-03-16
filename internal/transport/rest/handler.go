package rest

import (
	"github.com/dnevsky/restaurant-back/internal/pkg/config"
	"github.com/dnevsky/restaurant-back/internal/service"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/helpers"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/middleware"
	v1 "github.com/dnevsky/restaurant-back/internal/transport/rest/v1"
	"net/http"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	helpers  helpers.Helpers
}

func NewHandler(services *service.Service, helperManager *helpers.Manager) *Handler {
	return &Handler{
		services: services,
		helpers:  helperManager,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	if config.Config.Env == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	prometheus := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	router.Use(
		middleware.PanicRecovery(),
		middleware.Limit(config.Config.Limiter.RPS, config.Config.Limiter.Burst, config.Config.Limiter.TTL),
		middleware.Cors(),
		prometheus.Instrument(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.HEAD("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	if config.Config.PprofEnabled {
		pprof.Register(router)
	}

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	h.InitRoot(router)
	handlerV1 := v1.NewHandler(h.services, h.helpers)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
