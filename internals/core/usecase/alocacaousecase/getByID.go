package alocacaousecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) GetByID(id int32) (*domain.Alocacao, error) {
	alocacao, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return alocacao, nil
}