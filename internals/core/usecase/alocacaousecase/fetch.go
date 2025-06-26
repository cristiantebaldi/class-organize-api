package alocacaousecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) Fetch() (*[]domain.Alocacao, error) {
	alocacoes, err := usecase.repository.Fetch()
	if err != nil {
		return nil, err
	}

	return alocacoes, nil
}