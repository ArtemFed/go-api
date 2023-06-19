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
		students := api.Group("/students")
		{
			students.GET("/get-all", h.getAllStudents)
			students.GET("/get-by-id", h.getStudentById)
			students.POST("/create", h.createStudent)
			students.PUT("/update", h.updateStudent)
			students.DELETE("/delete", h.deleteGrade)
		}

		grades := api.Group("/grades")
		{
			grades.GET("/get-by-id", h.getGradeById)
			grades.GET("/get-by-student-id", h.getGradesByStudentId)
			grades.POST("/create", h.createGrade)
			grades.DELETE("/delete", h.deleteGrade)
		}
	}

	return router
}
