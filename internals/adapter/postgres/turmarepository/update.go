package turmarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Update(
	id int32,
	turmaRequest *dto.TurmaRequestBody,
) (*domain.Turma, error) {
	turma := domain.Turma{}

	err := repository.db.QueryRowx(`
		UPDATE turmas SET 
		nome = $1, 
		curso = $2, 
		periodo = $3, 
		quant_alunos = $4 
		WHERE id = $5
		returning *;
	`,
		turmaRequest.Nome,
		turmaRequest.Curso,
		turmaRequest.Periodo,
		turmaRequest.QuantAlunos,
		id,
	).StructScan(&turma)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &turma, nil
}
