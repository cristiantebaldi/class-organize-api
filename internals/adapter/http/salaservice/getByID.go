package salaservice

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (service service) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	sala, err := service.usecase.GetByID(int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, sala)
}
