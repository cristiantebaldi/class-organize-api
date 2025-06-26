package professorrepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Update(
	id int32,
	professorRequest *dto.ProfessorRequestBody,
) (*domain.Professor, error) {
	professor := domain.Professor{}

	err := repository.db.QueryRowx(`
		UPDATE professores SET 
		nome = $1, 
		email = $2, 
		formacao = $3, 
		disciplina = $4 
		WHERE id = $5 returning *;
	`,
		professorRequest.Nome,
		professorRequest.Email,
		professorRequest.Formacao,
		professorRequest.Disciplina,
		id,
	).StructScan(&professor)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &professor, nil
}
