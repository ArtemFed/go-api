package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/src/transport/models/payloads"
	"net/http"
)

func (h *Handler) getGradeById(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	grade, err := h.services.Grade.GetGradeById(id)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	response := payloads.ConvertGradeDTOToGradeResponse(grade)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) getGradesByStudentId(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	grades, err := h.services.Grade.GetGradesByStudentId(id)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	array := payloads.GradesArrayResponse{Grades: make([]payloads.GradeResponse, len(*grades))}
	for i, grade := range *grades {
		array.Grades[i] = *payloads.ConvertGradeDTOToGradeResponse(&grade)
	}
	c.JSON(http.StatusOK, array)
}

func (h *Handler) createGrade(c *gin.Context) {
	var input payloads.GradePayload
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.SubjectName) > 60 {
		newErrorResponse(c, http.StatusBadRequest, "too long string value for subject name")
		return
	}

	if input.Grade < 0 || input.Grade > 10 {
		newErrorResponse(c, http.StatusBadRequest, "invalid value for grade (0 < age <= 100)")
		return
	}

	gradeDTO := payloads.ConvertGradePayloadToGradeDTO(&input)

	student, err := h.services.Student.GetStudentById(gradeDTO.StudentId)
	if student == nil {
		newErrorResponse(c, http.StatusConflict, "student does noy exist")
		return
	}
	if err == nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.CreateGrade(gradeDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response": fmt.Sprintf("grade with id %d successfully created", id),
	})
}

func (h *Handler) deleteGrade(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	grade, err := h.services.Grade.GetGradeById(id)
	if grade == nil {
		newErrorResponse(c, http.StatusConflict, "grade does not exist")
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteGrade(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response": fmt.Sprintf("grade with id %d successfully deleted", id),
	})
}
