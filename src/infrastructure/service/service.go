package service

import (
	"go-api/src/domain"
	"go-api/src/infrastructure/repository"
)

type Material interface {
	GetMaterialById(materialId int) (*domain.MaterialDTO, error)
	GetAllMaterials() ([]domain.MaterialDTO, error)
	CreateMaterial(material *domain.MaterialDTO) (int, error)
	UpdateMaterial(material *domain.MaterialDTO) (int, error)
	DeleteMaterial(materialId int) error
}

type Service struct {
	Material
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Material: NewMaterialService(repos.Material),
	}
}
