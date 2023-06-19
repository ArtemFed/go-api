package payloads

import (
	"go-api/src/domain"
)

type StudentDTO struct {
	Id        int
	Name      string
	Age       int
	Specialty string
}

type StudentPayload struct {
	Id        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Specialty string `json:"specialty" binding:"required"`
}

type StudentResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Specialty string `json:"specialty"`
}

type StudentsArrayResponse struct {
	Students []StudentResponse `json:"students"`
}

func ConvertStudentPayloadToStudentDTO(payload *StudentPayload) *domain.StudentDTO {
	return &domain.StudentDTO{
		Id:        payload.Id,
		Name:      payload.Name,
		Age:       payload.Age,
		Specialty: payload.Specialty,
	}
}

func ConvertStudentDTOToStudentResponse(dto *domain.StudentDTO) *StudentResponse {
	return &StudentResponse{
		Id:        dto.Id,
		Name:      dto.Name,
		Age:       dto.Age,
		Specialty: dto.Specialty,
	}
}
