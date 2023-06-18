package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func getContextId(c *gin.Context) (int, error) {
	id, ok := c.GetQuery("id")
	if !ok {
		return 0, errors.New("id not found" + id)
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("id is of invalid type")
	}

	log.Printf("func getContextId(): id=%d", id)
	return idInt, nil
}
