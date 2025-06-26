package domain

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

type Sala struct {
	ID         int    `json:"id"`
	Numero     string `json:"numero"`
	Capacidade int    `json:"capacidade"`
	Bloco      string `json:"bloco"`
	Tipo       string `json:"tipo"`
}

type SalaHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
	Fetch(*gin.Context)
}

type SalaUsecase interface {
	Create(*dto.SalaRequestBody) (*Sala, error)
	Update(int32, *dto.SalaRequestBody) (*Sala, error)
	Delete(int32) (*Sala, error)
	GetByID(int32) (*Sala, error)
	Fetch() (*[]Sala, error)
}

type SalaRepository interface {
	Create(*dto.SalaRequestBody) (*Sala, error)
	Update(int32, *dto.SalaRequestBody) (*Sala, error)
	Delete(int32) (*Sala, error)
	GetByID(int32) (*Sala, error)
	Fetch() (*[]Sala, error)
}
