package turmaservice

import (
	"net/http"

	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	turmaRequest, err := dto.FromJsonTurmaRequestBody(c.Request.Body, service.validator)

	if err != nil {
		c.JSON(http.StatusInternalServerError, turmaRequest)
		return
	}

	turma, err := service.usecase.Create(turmaRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, turma)
}
