package alocacaorepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) GetSalasAlocadas(diaSemana string, horarioInicio string, horarioFim string) (*[]domain.Alocacao, error) {
	alocacoes := []domain.Alocacao{}

	query := `
		SELECT DISTINCT sala_id FROM alocacoes 
		WHERE dia_semana = $1 AND 
		((horario_inicio <= $2 AND horario_fim > $2) OR 
		(horario_inicio < $3 AND horario_fim >= $3) OR
		(horario_inicio >= $2 AND horario_fim <= $3))
	`
	{
		err := repository.db.Select(&alocacoes, query, diaSemana, horarioInicio, horarioFim)

		if err != nil {
			return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
		}
	}

	return &alocacoes, nil
}