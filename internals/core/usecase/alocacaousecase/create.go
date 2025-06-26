package alocacaousecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Create(alocacaoRequest *dto.AlocacaoRequestBody) (*domain.Alocacao, error) {
	alocacao, err := usecase.repository.Create(alocacaoRequest)

	if err != nil {
		return nil, err
	}

	return alocacao, nil
}
