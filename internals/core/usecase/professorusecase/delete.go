package professorusecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) Delete(id int32) (*domain.Professor, error) {
	professor, err := usecase.repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return professor, nil
}
