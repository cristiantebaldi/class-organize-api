package professorservice

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

	professorRequest, err := dto.FromJsonProfessorRequestBody(c.Request.Body, service.validator)

	if err != nil {
		c.JSON(http.StatusInternalServerError, professorRequest)
		return
	}

	professor, err := service.usecase.Update(int32(id), professorRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, professor)
}
