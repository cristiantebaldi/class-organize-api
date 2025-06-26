package salausecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Update(
	id int32,
	salaRequest *dto.SalaRequestBody,
) (*domain.Sala, error) {
	sala, err := usecase.repository.Update(id, salaRequest)

	if err != nil {
		return nil, err
	}

	return sala, nil
}
