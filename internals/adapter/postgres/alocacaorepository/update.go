package alocacaorepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Update(
	id int32,
	alocacaoRequest *dto.AlocacaoRequestBody,
) (*domain.Alocacao, error) {
	alocacao := domain.Alocacao{}

	err := repository.db.QueryRowx(`
		UPDATE alocacoes SET 
			professor_id = $1, 
			sala_id = $2, 
			turma_id = $3, 
			dia_semana = $4, 
			horario_inicio = $5, 
			horario_fim = $6 
		WHERE id = $7 returning *;
	`,
		alocacaoRequest.ProfessorID,
		alocacaoRequest.SalaID,
		alocacaoRequest.TurmaID,
		alocacaoRequest.DiaSemana,
		alocacaoRequest.HorarioInicio,
		alocacaoRequest.HorarioFim,
		id,
	).StructScan(&alocacao)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &alocacao, nil
}
