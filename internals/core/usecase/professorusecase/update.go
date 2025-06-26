package professorusecase

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) Update(
	id int32,
	professorRequest *dto.ProfessorRequestBody,
) (*domain.Professor, error) {
	professor, err := usecase.repository.Update(id, professorRequest)

	if err != nil {
		return nil, err
	}

	return professor, nil
}
