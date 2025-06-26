package alocacaorepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Fetch() (*[]domain.Alocacao, error) {
	alocacoes := []domain.Alocacao{}

	query := `SELECT 
			a.id, a.professor_id, a.sala_id, a.turma_id, a.dia_semana, a.horario_inicio, a.horario_fim,
			p.id, p.nome, p.email, p.formacao, p.disciplina,
			s.id, s.numero, s.capacidade, s.bloco, s.tipo,
			t.id, t.nome AS turma_nome, t.curso, t.periodo, t.quant_alunos
		FROM alocacoes a
		JOIN professores p ON a.professor_id = p.id
		JOIN salas s ON a.sala_id = s.id
		JOIN turmas t ON a.turma_id = t.id`
	{
		err := repository.db.Select(&alocacoes, query)

		if err != nil {
			return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
		}
	}

	return &alocacoes, nil
}
