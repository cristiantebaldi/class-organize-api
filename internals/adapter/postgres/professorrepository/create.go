package professorrepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Create(
	professorRequest *dto.ProfessorRequestBody,
) (*domain.Professor, error) {
	professor := domain.Professor{}

	err := repository.db.QueryRowx(`
		INSERT INTO professores (nome, formacao, email, disciplina) 
		VALUES ($1, $2, $3, $4) returning *;
	`,
		professorRequest.Nome,
		professorRequest.Formacao,
		professorRequest.Email,
		professorRequest.Disciplina,
	).StructScan(&professor)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &professor, nil
}