package repository

import (
	"github.com/jmoiron/sqlx"
	"go-api/src/domain"
)

type Student interface {
	GetStudentById(studentId int) (*domain.StudentDTO, error)
	GetAllStudents() (*[]domain.StudentDTO, error)
	CreateStudent(student *domain.StudentDTO) (int, error)
	UpdateStudent(student *domain.StudentDTO) (int, error)
	DeleteStudent(studentId int) error
}

type Grade interface {
	GetGradeById(gradeId int) (*domain.GradeDTO, error)
	GetGradesByStudentId(gradeId int) (*[]domain.GradeDTO, error)
	CreateGrade(grade *domain.GradeDTO) (int, error)
	DeleteGrade(gradeId int) error
}

type Repository struct {
	Student
	Grade
}

func NewRepository(
	db *sqlx.DB,
) *Repository {
	return &Repository{
		Student: NewStudentPostgres(db),
		Grade:   NewGradePostgres(db),
	}
}
