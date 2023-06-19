package payloads

import (
	"go-api/src/domain"
	"time"
)

type GradePayload struct {
	Id          int    `json:"id"`
	Grade       int    `json:"grade" binding:"required"`
	StudentId   int    `json:"student_id" binding:"required"`
	SubjectName string `json:"subject_name" binding:"required"`
}

type GradeResponse struct {
	Id          int       `json:"id"`
	Grade       int       `json:"grade"`
	StudentId   int       `json:"student_id"`
	SubjectName string    `json:"subject_name"`
	PublishDate time.Time `json:"publish_date"`
}

type GradesArrayResponse struct {
	Grades []GradeResponse `json:"grades"`
}

func ConvertGradePayloadToGradeDTO(payload *GradePayload) *domain.GradeDTO {
	return &domain.GradeDTO{
		Id:          payload.Id,
		Grade:       payload.Grade,
		StudentId:   payload.StudentId,
		SubjectName: payload.SubjectName,
	}
}

func ConvertGradeDTOToGradeResponse(dto *domain.GradeDTO) *GradeResponse {
	return &GradeResponse{
		Id:          dto.Id,
		Grade:       dto.Grade,
		StudentId:   dto.StudentId,
		SubjectName: dto.SubjectName,
		PublishDate: dto.PublishDate,
	}
}
