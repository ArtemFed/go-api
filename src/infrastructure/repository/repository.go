package repository

import (
	"github.com/jmoiron/sqlx"
	"go-api/src/domain"
)

type Material interface {
	GetMaterialById(materialId int) (*domain.MaterialDTO, error)
	GetAllMaterials() ([]domain.MaterialDTO, error)
	CreateMaterial(material *domain.MaterialDTO) (int, error)
	UpdateMaterial(material *domain.MaterialDTO) (int, error)
	DeleteMaterial(materialId int) error
}

type Repository struct {
	Material
}

func NewRepository(
	db *sqlx.DB,
) *Repository {
	return &Repository{
		Material: NewMaterialPostgres(db),
	}
}
