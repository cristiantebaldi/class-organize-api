package professorusecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

type usecase struct {
	repository domain.ProfessorRepository
}

func New(repository domain.ProfessorRepository) domain.ProfessorUsecase {
	return &usecase{
		repository: repository,
	}
}