package domain

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

type Turma struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Curso       string `json:"curso"`
	Periodo     string `json:"periodo"`
	QuantAlunos int    `json:"quant_alunos" db:"quant_alunos"`
}

type TurmaHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
	Fetch(*gin.Context)
}

type TurmaUsecase interface {
	Create(*dto.TurmaRequestBody) (*Turma, error)
	Update(int32, *dto.TurmaRequestBody) (*Turma, error)
	Delete(int32) (*Turma, error)
	GetByID(int32) (*Turma, error)
	Fetch() (*[]Turma, error)
}

type TurmaRepository interface {
	Create(*dto.TurmaRequestBody) (*Turma, error)
	Update(int32, *dto.TurmaRequestBody) (*Turma, error)
	Delete(int32) (*Turma, error)
	GetByID(int32) (*Turma, error)
	Fetch() (*[]Turma, error)
}
