package turmausecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Update(
	id int32,
	turmaRequest *dto.TurmaRequestBody,
) (*domain.Turma, error) {
	turma, err := usecase.repository.Update(id, turmaRequest)

	if err != nil {
		return nil, err
	}

	return turma, nil
}
