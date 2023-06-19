package models

import (
	"go-api/src/domain"
	"time"
)

type GradeTable struct {
	Id          int       `db:"id"`
	Grade       int       `db:"grade"`
	StudentId   int       `db:"student_id"`
	SubjectName string    `db:"subject_name"`
	PublishDate time.Time `db:"publish_date"`
}

func ConvertGradeDTOToGradeTable(dto *domain.GradeDTO) *GradeTable {
	return &GradeTable{
		Id:          dto.Id,
		Grade:       dto.Grade,
		StudentId:   dto.StudentId,
		SubjectName: dto.SubjectName,
		PublishDate: dto.PublishDate,
	}
}
func ConvertGradeTableToGradeDTO(table *GradeTable) *domain.GradeDTO {
	return &domain.GradeDTO{
		Id:          table.Id,
		Grade:       table.Grade,
		StudentId:   table.StudentId,
		SubjectName: table.SubjectName,
		PublishDate: table.PublishDate,
	}
}
