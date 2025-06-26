package turmausecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Create(turmaRequest *dto.TurmaRequestBody) (*domain.Turma, error) {
	turma, err := usecase.repository.Create(turmaRequest)

	if err != nil {
		return nil, err
	}

	return turma, nil
}
