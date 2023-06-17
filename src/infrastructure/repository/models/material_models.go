package models

import (
	"github.com/north-hascii/crm-planning/planning_service/src/domain/models"
)

type MaterialTable struct {
	Id           int    `db:"id"`
	MaterialName string `db:"material_name"`
	UnitName     string `db:"units"`
}

func ConvertMaterialDTOToMaterialTable(dto *models.MaterialDTO) *MaterialTable {
	return &MaterialTable{
		MaterialName: dto.MaterialName,
		UnitName:     dto.UnitName,
	}
}
func ConvertMaterialTableToMaterialDTO(table *MaterialTable) *models.MaterialDTO {
	return &models.MaterialDTO{
		Id:           table.Id,
		MaterialName: table.MaterialName,
		UnitName:     table.UnitName,
	}
}
