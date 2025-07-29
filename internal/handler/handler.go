package handler

import (
	"vk-server-task/internal/metrics"
	"vk-server-task/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(metrics.GinPrometheusMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	api := router.Group("/", h.AuthMiddleware)
	{
		api.POST("/ads", h.Create)
	}

	router.GET("/ads", h.Get)

	return router
}
