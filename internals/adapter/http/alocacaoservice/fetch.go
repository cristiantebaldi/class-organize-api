package alocacaoservice

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) Fetch(c *gin.Context) {
	alocacoes, err := service.usecase.Fetch()

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, alocacoes)
}
