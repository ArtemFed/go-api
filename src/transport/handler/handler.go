package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/src/infrastructure/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		material := api.Group("/material")
		{
			material.GET("/get-all", h.getAllMaterials)
			material.GET("/get-by-id", h.getMaterialById)
			material.POST("/create", h.createMaterial)
			material.PUT("/update", h.updateMaterial)
			material.DELETE("/delete", h.deleteMaterial)
		}
	}

	return router
}
