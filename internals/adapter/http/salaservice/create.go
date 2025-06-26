package salaservice

import (
	"net/http"

	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	salaRequest, err := dto.FromJsonSalaRequestBody(c.Request.Body, service.validator)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, salaRequest)
		return
	}

	sala, err := service.usecase.Create(salaRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, sala)
}
