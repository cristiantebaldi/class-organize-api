package turmausecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

type usecase struct {
	repository domain.TurmaRepository
}

func New(repository domain.TurmaRepository) domain.TurmaUsecase {
	return &usecase{
		repository: repository,
	}
}