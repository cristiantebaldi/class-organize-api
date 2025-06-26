package alocacaorepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Create(
	alocacaoRequest *dto.AlocacaoRequestBody,
) (*domain.Alocacao, error) {
	alocacao := domain.Alocacao{}

	err := repository.db.QueryRowx(`
		INSERT INTO alocacoes (professor_id, sala_id, turma_id, dia_semana, horario_inicio, horario_fim) 
		VALUES ($1, $2, $3, $4, $5, $6) returning *;
	`,
		alocacaoRequest.ProfessorID,
		alocacaoRequest.SalaID,
		alocacaoRequest.TurmaID,
		alocacaoRequest.DiaSemana,
		alocacaoRequest.HorarioInicio,
		alocacaoRequest.HorarioFim,
	).StructScan(&alocacao)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &alocacao, nil
}