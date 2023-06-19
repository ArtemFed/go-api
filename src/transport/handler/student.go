package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/src/transport/models/payloads"
	"net/http"
)

func (h *Handler) getAllStudents(c *gin.Context) {
	students, err := h.services.Student.GetAllStudents()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	array := payloads.StudentsArrayResponse{Students: make([]payloads.StudentResponse, len(*students))}
	for i, student := range *students {
		array.Students[i] = *payloads.ConvertStudentDTOToStudentResponse(&student)
	}
	c.JSON(http.StatusOK, array)
}

func (h *Handler) getStudentById(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	student, err := h.services.Student.GetStudentById(id)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	response := payloads.ConvertStudentDTOToStudentResponse(student)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) createStudent(c *gin.Context) {
	var input payloads.StudentPayload
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.Name) > 60 {
		newErrorResponse(c, http.StatusBadRequest, "too long string value for student name")
		return
	}

	if input.Age < 0 || input.Age > 100 {
		newErrorResponse(c, http.StatusBadRequest, "invalid value for age (0 < age <= 100)")
		return
	}

	studentDTO := payloads.ConvertStudentPayloadToStudentDTO(&input)

	id, err := h.services.CreateStudent(studentDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response": fmt.Sprintf("student with id %d successfully created", id),
	})
}

func (h *Handler) updateStudent(c *gin.Context) {
	var input payloads.StudentPayload
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.Name) > 60 {
		newErrorResponse(c, http.StatusBadRequest, "too long string value for student name")
		return
	}

	if input.Age < 0 || input.Age > 100 {
		newErrorResponse(c, http.StatusBadRequest, "invalid value for age (0 < age <= 100)")
		return
	}

	studentDTO := payloads.ConvertStudentPayloadToStudentDTO(&input)

	student, err := h.services.Student.GetStudentById(studentDTO.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if student == nil {
		newErrorResponse(c, http.StatusConflict, "student does not exist")
		return
	}

	id, err := h.services.UpdateStudent(studentDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response": fmt.Sprintf("student with id %d successfully updated", id),
	})
}

func (h *Handler) deleteStudent(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	student, err := h.services.Student.GetStudentById(id)
	if student == nil {
		newErrorResponse(c, http.StatusConflict, "student does not exist")
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteStudent(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response": fmt.Sprintf("student with id %d successfully deleted", id),
	})
}
