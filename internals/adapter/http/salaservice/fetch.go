package salaservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) Fetch(c *gin.Context) {
	salas, err := service.usecase.Fetch()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, salas)
}
