package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/src/transport/models/payloads"
	"net/http"
	"strconv"
)

func (h *Handler) getAllMaterials(c *gin.Context) {
	materials, err := h.services.Material.GetAllMaterials()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	array := payloads.MaterialsArrayResponse{Materials: make([]payloads.MaterialResponse, len(materials))}
	for i, material := range materials {
		array.Materials[i] = *payloads.ConvertMaterialDTOToMaterialResponse(&material)
	}
	c.JSON(http.StatusOK, array)
}

func (h *Handler) getMaterialById(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	material, err := h.services.Material.GetMaterialById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, payloads.ConvertMaterialDTOToMaterialResponse(material))
}

func (h *Handler) createMaterial(c *gin.Context) {
	var input payloads.MaterialPayload
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.MaterialName) > 60 ||
		len(input.UnitName) > 30 {
		newErrorResponse(c, http.StatusBadRequest, "too long string value")
		return
	}

	materialDTO := payloads.ConvertMaterialPayloadToMaterialDTO(&input)
	id, err := h.services.CreateMaterial(materialDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"request": "create material id " + strconv.Itoa(id),
	})
}

func (h *Handler) updateMaterial(c *gin.Context) {
	var input payloads.MaterialPayload
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.MaterialName) > 60 ||
		len(input.UnitName) > 30 {
		newErrorResponse(c, http.StatusBadRequest, "too long string value")
		return
	}

	materialDTO := payloads.ConvertMaterialPayloadToMaterialDTO(&input)
	id, err := h.services.UpdateMaterial(materialDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"request": "update material id:" + strconv.Itoa(id),
	})
}

func (h *Handler) deleteMaterial(c *gin.Context) {
	id, err := getContextId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	material, err := h.services.GetMaterialById(id)
	if material == nil || err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	err = h.services.DeleteMaterial(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"material_id": id,
		"request":     "delete material",
	})
}
