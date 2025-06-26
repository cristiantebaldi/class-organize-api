package salausecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) GetByID(id int32) (*domain.Sala, error) {
	sala, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return sala, nil
}
