package salausecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) Fetch() (*[]domain.Sala, error) {
	salas, err := usecase.repository.Fetch()
	if err != nil {
		return nil, err
	}

	return salas, nil
}
