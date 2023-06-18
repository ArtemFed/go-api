package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/src/domain"
	"go-api/src/infrastructure/repository/models"
	"log"
)

type MaterialPostgres struct {
	db *sqlx.DB
}

func NewMaterialPostgres(db *sqlx.DB) *MaterialPostgres {
	return &MaterialPostgres{db: db}
}

// GetMaterialById Получить по ID
func (p *MaterialPostgres) GetMaterialById(materialId int) (*domain.MaterialDTO, error) {
	var material models.MaterialTable
	var query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", MaterialsTable)
	err := p.db.Get(&material, query, materialId)
	if err != nil {
		log.Printf("Level: repos; func GetMaterialById(): material with name = %d does not exist",
			materialId,
		)
		return nil, err
	}

	materialDto := models.ConvertMaterialTableToMaterialDTO(&material)
	return materialDto, err
}

// GetAllMaterials Получить все
func (p *MaterialPostgres) GetAllMaterials() ([]domain.MaterialDTO, error) {
	var materials []models.MaterialTable
	var query = fmt.Sprintf("SELECT * FROM %s ORDER BY id", MaterialsTable)
	err := p.db.Select(&materials, query)
	log.Printf("Level: repos; func GetAllMaterials(): []domain.MaterialDTO=%v", materials)

	materialDTOs := make([]domain.MaterialDTO, len(materials))
	for i, material := range materials {
		materialDTOs[i] = *models.ConvertMaterialTableToMaterialDTO(&material)
	}
	return materialDTOs, err
}

// CreateMaterial Создать
func (p *MaterialPostgres) CreateMaterial(material *domain.MaterialDTO) (int, error) {
	var id int
	var createMaterialQuery = fmt.Sprintf(
		"INSERT INTO %s (material_name, units) VALUES ($1, $2) RETURNING id",
		MaterialsTable)
	row := p.db.QueryRow(createMaterialQuery,
		material.MaterialName,
		material.UnitName,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	log.Printf("Level: repos; func createMaterial(): domain.MaterialDTO=%v", material)
	return id, nil
}

// UpdateMaterial Обновить данные
func (p *MaterialPostgres) UpdateMaterial(material *domain.MaterialDTO) (int, error) {
	var updateMaterialQuery = fmt.Sprintf(
		"UPDATE %s SET material_name = $1, units = $2 WHERE id = $3",
		MaterialsTable,
	)
	_, err := p.db.Exec(updateMaterialQuery,
		material.MaterialName,
		material.UnitName,
		material.Id,
	)
	if err != nil {
		return 0, err
	}

	log.Printf("Level: repos; func UpdateMaterial(): domain.MaterialDTO=%v", material)
	return material.Id, nil
}

// DeleteMaterial Удалить
func (p *MaterialPostgres) DeleteMaterial(materialId int) error {
	var deleteMaterialQuery = fmt.Sprintf("DELETE FROM %s WHERE id = $1", MaterialsTable)
	_, err := p.db.Exec(deleteMaterialQuery,
		materialId,
	)
	if err != nil {
		return err
	}

	log.Printf("Level: repos; func DeleteMaterial(): Material id=%d", materialId)
	return nil
}
