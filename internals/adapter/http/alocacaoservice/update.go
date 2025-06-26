package alocacaoservice

import (
	"net/http"
	"strconv"

	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

func (service service) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	alocacaoRequest, err := dto.FromJsonAlocacaoRequestBody(c.Request.Body, service.validator)

	if err != nil {
		c.JSON(http.StatusInternalServerError, alocacaoRequest)
		return
	}

	alocacao, err := service.usecase.Update(int32(id), alocacaoRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, alocacao)
}
