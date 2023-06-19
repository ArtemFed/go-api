package service

import (
	"go-api/src/domain"
	"go-api/src/infrastructure/repository"
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
	GetGradesByStudentId(studentId int) (*[]domain.GradeDTO, error)
	CreateGrade(grade *domain.GradeDTO) (int, error)
	DeleteGrade(gradeId int) error
}

type Service struct {
	Student
	Grade
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Student: NewStudentService(repos.Student),
		Grade:   NewGradeService(repos.Grade),
	}
}
