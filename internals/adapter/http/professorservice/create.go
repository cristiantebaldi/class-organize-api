package professorservice

import (
	"net/http"

	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	professorRequest, err := dto.FromJsonProfessorRequestBody(c.Request.Body, service.validator)

	if err != nil {
		c.JSON(http.StatusInternalServerError, professorRequest)
		return
	}

	professor, err := service.usecase.Create(professorRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, professor)
}
