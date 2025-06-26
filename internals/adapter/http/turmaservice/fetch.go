package turmaservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) Fetch(c *gin.Context) {
	turmas, err := service.usecase.Fetch()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, turmas)
}
