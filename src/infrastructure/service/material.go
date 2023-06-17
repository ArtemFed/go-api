package service

import (
	"go-api/src/domain"
	"go-api/src/infrastructure/repository"
)

type MaterialService struct {
	repo repository.Material
}

func NewMaterialService(repo repository.Material) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) GetMaterialById(materialId int) (*domain.MaterialDTO, error) {
	return s.repo.GetMaterialById(materialId)
}

func (s *MaterialService) GetAllMaterials() ([]domain.MaterialDTO, error) {
	return s.repo.GetAllMaterials()
}

func (s *MaterialService) CreateMaterial(material *domain.MaterialDTO) (int, error) {
	return s.repo.CreateMaterial(material)
}

func (s *MaterialService) UpdateMaterial(material *domain.MaterialDTO) (int, error) {
	return s.repo.UpdateMaterial(material)
}

func (s *MaterialService) DeleteMaterial(materialId int) error {
	return s.repo.DeleteMaterial(materialId)
}
