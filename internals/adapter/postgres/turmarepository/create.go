package turmarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Create(
	turmaRequest *dto.TurmaRequestBody,
) (*domain.Turma, error) {
	turma := domain.Turma{}

	err := repository.db.QueryRowx(`
		INSERT INTO turmas (nome, curso, periodo, quant_alunos) 
		VALUES ($1, $2, $3, $4) returning *;
	`,
		turmaRequest.Nome,
		turmaRequest.Curso,
		turmaRequest.Periodo,
		turmaRequest.QuantAlunos,
	).StructScan(&turma)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &turma, nil
}