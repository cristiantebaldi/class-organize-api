package professorusecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Create(professorRequest *dto.ProfessorRequestBody) (*domain.Professor, error) {
	professor, err := usecase.repository.Create(professorRequest)

	if err != nil {
		return nil, err
	}

	return professor, nil
}
