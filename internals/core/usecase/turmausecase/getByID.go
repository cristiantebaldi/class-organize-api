package turmausecase

import "github.com/cristiantebaldi/class-organize-api/internals/core/domain"

func (usecase usecase) GetByID(id int32) (*domain.Turma, error) {
	turma, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return turma, nil
}
