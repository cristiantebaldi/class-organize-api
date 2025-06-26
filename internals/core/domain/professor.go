package domain

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

type Professor struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	Formacao   string `json:"formacao"`
	Disciplina string `json:"disciplina"`
}

type ProfessorHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
	Fetch(*gin.Context)
}

type ProfessorUsecase interface {
	Create(*dto.ProfessorRequestBody) (*Professor, error)
	Update(int32, *dto.ProfessorRequestBody) (*Professor, error)
	Delete(int32) (*Professor, error)
	GetByID(int32) (*Professor, error)
	Fetch() (*[]Professor, error)
}

type ProfessorRepository interface {
	Create(*dto.ProfessorRequestBody) (*Professor, error)
	Update(int32, *dto.ProfessorRequestBody) (*Professor, error)
	Delete(int32) (*Professor, error)
	GetByID(int32) (*Professor, error)
	Fetch() (*[]Professor, error)
}
