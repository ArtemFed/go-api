package payloads

import (
	"github.com/north-hascii/crm-planning/planning_service/src/domain/models"
)

type MaterialPayload struct {
	Id           int    `json:"id"`
	MaterialName string `json:"material_name" binding:"required"`
	UnitName     string `json:"units" binding:"required"`
}

type MaterialResponse struct {
	Id           int    `json:"id"`
	MaterialName string `json:"material_name" binding:"required"`
	UnitName     string `json:"units" binding:"required"`
}

type MaterialsArrayResponse struct {
	Materials []MaterialResponse `json:"materials"`
}

func ConvertMaterialPayloadToMaterialDTO(payload *MaterialPayload) *models.MaterialDTO {
	return &models.MaterialDTO{
		Id:           payload.Id,
		MaterialName: payload.MaterialName,
		UnitName:     payload.UnitName,
	}
}

func ConvertMaterialDTOToMaterialResponse(dto *models.MaterialDTO) *MaterialResponse {
	return &MaterialResponse{
		Id:           dto.Id,
		MaterialName: dto.MaterialName,
		UnitName:     dto.UnitName,
	}
}
