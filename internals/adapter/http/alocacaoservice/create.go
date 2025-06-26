package alocacaoservice

import (
	"net/http"

	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	alocacaoRequest, err := dto.FromJsonAlocacaoRequestBody(c.Request.Body, service.validator)

	if err != nil {
		c.JSON(http.StatusInternalServerError, alocacaoRequest)
		return
	}

	alocacao, err := service.usecase.Create(alocacaoRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, alocacao)
}
