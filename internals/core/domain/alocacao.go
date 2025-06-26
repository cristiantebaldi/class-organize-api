package domain

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
	"github.com/gin-gonic/gin"
)

type Alocacao struct {
	ID                  int     `json:"id"`
	ProfessorID         int     `json:"professor_id" db:"professor_id"`
	SalaID              int     `json:"sala_id" db:"sala_id"`
	TurmaID             int     `json:"turma_id" db:"turma_id"`
	DiaSemana           string  `json:"dia_semana" db:"dia_semana"`
	HorarioInicio       string  `json:"horario_inicio" db:"horario_inicio"`
	HorarioFim          string  `json:"horario_fim" db:"horario_fim"`
	ProfessorNome       *string `json:"ProfessorNome" db:"nome"`
	ProfessorEmail      *string `json:"ProfessorEmail" db:"email"`
	ProfessorFormacacao *string `json:"ProfessorFormacacao" db:"formacao"`
	ProfessorDisciplina *string `json:"ProfessorDisciplina" db:"disciplina"`
	SalaNumero          *string `json:"SalaNumero" db:"numero"`
	SalaCapacidade      *int    `json:"SalaCapacidade" db:"capacidade"`
	SalaBloco           *string `json:"SalaBloco" db:"bloco"`
	SalaTipo            *string `json:"SalaTipo" db:"tipo"`
	TurmaNome           *string `json:"TurmaNome" db:"turma_nome"`
	TurmaCurso          *string `json:"TurmaCurso" db:"curso"`
	TurmaPeriodo        *string `json:"TurmaPeriodo" db:"periodo"`
	TurmaQuantAlunos    *int    `json:"TurmaQuantAlunos" db:"quant_alunos"`
}

type AlocacaoHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
	Fetch(*gin.Context)
	CreateAlocacoesAutomaticamente(*gin.Context)
}

type AlocacaoUsecase interface {
	Create(*dto.AlocacaoRequestBody) (*Alocacao, error)
	Update(int32, *dto.AlocacaoRequestBody) (*Alocacao, error)
	Delete(int32) (*Alocacao, error)
	GetByID(int32) (*Alocacao, error)
	Fetch() (*[]Alocacao, error)
	CreateAlocacoesAutomaticamente(*dto.AlocacaoRequestBody) (*[]Alocacao, error)
}

type AlocacaoRepository interface {
	Create(*dto.AlocacaoRequestBody) (*Alocacao, error)
	Update(int32, *dto.AlocacaoRequestBody) (*Alocacao, error)
	Delete(int32) (*Alocacao, error)
	GetByID(int32) (*Alocacao, error)
	Fetch() (*[]Alocacao, error)
	GetSalasAlocadas(string, string, string) (*[]Alocacao, error)
	GetTurmasAlocadas(string, string, string) (*[]Alocacao, error)
	GetProfessoresAlocados(string, string, string) (*[]Alocacao, error)
}
