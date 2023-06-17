package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const (
	ctxId = "id"
)

func getContextId(c *gin.Context) (int, error) {
	id, ok := c.GetQuery(ctxId)
	if !ok {
		log.Printf("Level: middleware; func getContextId(): id not found")
		return 0, errors.New("id not found" + id)
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Level: middleware; func getContextId(): id is of invalid type")
		return 0, errors.New("id is of invalid type")
	}

	log.Printf("Level: middleware; func getContextId(): id=%d", id)
	return idInt, nil
}
