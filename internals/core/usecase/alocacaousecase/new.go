package alocacaousecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

type usecase struct {
	repository          domain.AlocacaoRepository
	professorRepository domain.ProfessorRepository
	salaRepository      domain.SalaRepository
	turmaRepository     domain.TurmaRepository
}

func New(
	repository domain.AlocacaoRepository,
	professorRepository domain.ProfessorRepository,
	salaRepository domain.SalaRepository,
	turmaRepository domain.TurmaRepository,
) domain.AlocacaoUsecase {
	return &usecase{
		repository:     repository,
		professorRepository: professorRepository,
		salaRepository: salaRepository,
		turmaRepository: turmaRepository,
	}
}
