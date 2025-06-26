package salausecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Create(salaRequest *dto.SalaRequestBody) (*domain.Sala, error) {
	sala, err := usecase.repository.Create(salaRequest)

	if err != nil {
		return nil, err
	}

	return sala, nil
}
