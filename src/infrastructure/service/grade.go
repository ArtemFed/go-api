package service

import (
	"go-api/src/domain"
	"go-api/src/infrastructure/repository"
)

type GradeService struct {
	repo repository.Grade
}

func NewGradeService(repo repository.Grade) *GradeService {
	return &GradeService{repo: repo}
}

func (s *GradeService) GetGradeById(gradeId int) (*domain.GradeDTO, error) {
	return s.repo.GetGradeById(gradeId)
}

func (s *GradeService) GetGradesByStudentId(studentId int) (*[]domain.GradeDTO, error) {
	return s.repo.GetGradesByStudentId(studentId)
}

func (s *GradeService) CreateGrade(grade *domain.GradeDTO) (int, error) {
	return s.repo.CreateGrade(grade)
}

func (s *GradeService) DeleteGrade(gradeId int) error {
	return s.repo.DeleteGrade(gradeId)
}
