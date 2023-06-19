package models

import (
	"go-api/src/domain"
)

type StudentTable struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Age       int    `db:"age"`
	Specialty string `db:"specialty"`
}

func ConvertStudentDTOToStudentTable(dto *domain.StudentDTO) *StudentTable {
	return &StudentTable{
		Id:        dto.Id,
		Name:      dto.Name,
		Age:       dto.Age,
		Specialty: dto.Specialty,
	}
}
func ConvertStudentTableToStudentDTO(table *StudentTable) *domain.StudentDTO {
	return &domain.StudentDTO{
		Id:        table.Id,
		Name:      table.Name,
		Age:       table.Age,
		Specialty: table.Specialty,
	}
}
