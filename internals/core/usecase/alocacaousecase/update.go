package alocacaousecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Update(
	id int32,
	alocacaoRequest *dto.AlocacaoRequestBody,
) (*domain.Alocacao, error) {
	alocacao, err := usecase.repository.Update(id, alocacaoRequest)

	if err != nil {
		return nil, err
	}

	return alocacao, nil
}
