package service

import (
	"go-api/src/domain"
	"go-api/src/infrastructure/repository"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetStudentById(studentId int) (*domain.StudentDTO, error) {
	return s.repo.GetStudentById(studentId)
}

func (s *StudentService) GetAllStudents() (*[]domain.StudentDTO, error) {
	return s.repo.GetAllStudents()
}

func (s *StudentService) CreateStudent(student *domain.StudentDTO) (int, error) {
	return s.repo.CreateStudent(student)
}

func (s *StudentService) UpdateStudent(student *domain.StudentDTO) (int, error) {
	return s.repo.UpdateStudent(student)
}

func (s *StudentService) DeleteStudent(studentId int) error {
	return s.repo.DeleteStudent(studentId)
}
