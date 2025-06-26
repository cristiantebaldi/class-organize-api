package professorusecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) Fetch() (*[]domain.Professor, error) {
	professores, err := usecase.repository.Fetch()
	if err != nil {
		return nil, err
	}

	return professores, nil
}
