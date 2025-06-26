package salausecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

type usecase struct {
	repository domain.SalaRepository
}

func New(repository domain.SalaRepository) domain.SalaUsecase {
	return &usecase{
		repository: repository,
	}
}
